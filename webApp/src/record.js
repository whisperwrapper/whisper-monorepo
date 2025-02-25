const { transcribeLiveWeb, connectionTest, getSessionId } = require('./send-file.js')
const { LANGUAGES } = require('./languages.js')

let resultField
let transcriptReceived = []
let transcriptIter = 0
let audioChunks = []
let recordedChunks = []
let canRecord = true
let shouldRecord = false
let isSending = false
let audioContext;
let myAudioWorkletNode

const blob = new Blob([`
    class MyProcessor extends AudioWorkletProcessor {
        process(inputs, outputs, parameters) {
            const input = inputs[0];
            if (input && input.length > 0) {
                this.port.postMessage(input[0]);
            }
            return true;
        }
    }
    registerProcessor('my-processor', MyProcessor);
`], { type: 'application/javascript' }
)
const url = URL.createObjectURL(blob)


window.onload = async function () {
    connectionTest()
    resultField = document.getElementById("transcription_result")
    micButton = document.getElementById("mic_button")
    micButton.addEventListener("click", () => changeMicState(this))
}

async function changeMicState(microphone) {
    const recording_dots = document.getElementById("recording_dots");
    // const icon = microphone.querySelector('.material-symbols-outlined'); //TODO: commented icon here
    // icon.innerText = recording ? 'stop_circle' : 'mic';
    if (canRecord) {
        const languages = document.getElementById("choose_record_lang");
        console.log(LANGUAGES)
        let selectedLanguage = languages.options[languages.selectedIndex].text;
        if (LANGUAGES.hasOwnProperty(selectedLanguage)) {
            selectedLanguage = LANGUAGES[selectedLanguage]
        } else {
            selectedLanguage = ""
        }
        console.log(selectedLanguage)
        recording_dots.classList.toggle('hidden');
        transcriptReceived = []
        transcriptIter = 0
        audioContext = new AudioContext( { sampleRate: 44100 })
        await audioContext.audioWorklet.addModule(url)
        myAudioWorkletNode = new AudioWorkletNode(audioContext, 'my-processor')
        recordAndSend(audioContext, myAudioWorkletNode, selectedLanguage)
    } else if (shouldRecord) {
        recording_dots.classList.toggle('hidden');
        stopRecording(audioContext, myAudioWorkletNode)
    } else {
        console.log("Poprzednia sesja jeszcze sie transkrybuje, poczekaj chwilkę :3")
    }
  }


async function recordAndSend(audioContext, myAudioWorkletNode, selectedLanguage) {
    if (canRecord) {
        console.log("button pressed")
        const stream = await navigator.mediaDevices.getUserMedia({ audio: { channelCount: 1} })
        const source = audioContext.createMediaStreamSource(stream)
        source.connect(myAudioWorkletNode)
        myAudioWorkletNode.connect(audioContext.destination)
    
        myAudioWorkletNode.port.onmessage = (event) => {
            audioChunks.push(event.data)
        };

        shouldRecord = true
        canRecord = false
        console.log("Before metadata")
        let sessionId = await getSessionId()
        let metadata = {"session_id": sessionId}

        recordAndSegmentChunks = setInterval(() => {
            if (shouldRecord && audioChunks.length > 0) {
                recordedChunks.push(audioChunks.splice(0))
            } else {
                clearInterval(recordAndSegmentChunks)
            }
        }, 2000)

        sendingInterval = setInterval(async () => {
            if (!isSending && recordedChunks.length > 0) {
                isSending = true
                const currentChunk = recordedChunks.shift(0)
                console.log("Wymsyłam")
                response = await transcribeLiveWeb(currentChunk, selectedLanguage, metadata)
                printResult(response)
                isSending = false
            } else if (!shouldRecord && recordedChunks.length == 0 && !isSending) {
                console.log("Que?")
                canRecord = true
                clearInterval(sendingInterval)
                return
            }
        }, 2000)
    }
}

function stopRecording(audioContext, myAudioWorkletNode) {
    shouldRecord = false
    myAudioWorkletNode.disconnect()
    audioContext.close()
    console.log("Recording finished")
}


function printResult(response) {
    resultField.textContent = ""
    console.log(response.getText())
    console.log(response.getNewChunk())
    transcriptReceived[transcriptIter] = response.getText()
    for (const entry of transcriptReceived) {
        resultField.textContent += entry
    }
    if (response.getNewChunk() == true) {
        transcriptIter += 1
        transcriptReceived.push([])
    }
}