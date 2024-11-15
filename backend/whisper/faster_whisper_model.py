from faster_whisper import WhisperModel
from pathlib import Path
import uuid
import os

import preprocess

model_size = "tiny"

class FasterWhisperHandler():
    
    def __init__(
        self, 
        language='en',
        task='transcribe'
    ):
        super(FasterWhisperHandler, self).__init__()
        self.model = WhisperModel(model_size, device="cpu", compute_type="float32", cpu_threads=8)
        self.language = language
        self.task = task
        self.silence = 1.5


    def preprocessStreaming(self, data, previousAudio, seconds, record):
        data = previousAudio + data
        previousAudio = data
        path = preprocess.saveFile(data, record)
        isSilence = preprocess.detectSilence(path, self.silence)
        return path, isSilence, previousAudio, seconds
    

    def preprocessRequest(self, data, record):
        return preprocess.saveFile(data, record)
    

    def transcribe(self, data, previous=""):
       #  print(f'Zapisane: {previous}')
        segments, info = self.model.transcribe(str(data), beam_size=5, language='pl', initial_prompt=previous)
        print("Detected language '%s' with probability %f" % (info.language, info.language_probability))

        response = ""
        for segment in segments:
            # print("[%.2fs -> %.2fs] %s" % (segment.start, segment.end, segment.text))
            response += segment.text
        data.unlink()
        # print(response)
        return response
    

    def postprocess(self, transcription, isSilence, segment, previousAudio, data, seconds):
        if (len(data) >= 3 and data[-3:] == '...'):
            data = data[0:-3]
        seconds += 2
        if isSilence or seconds >= 10:
            print("10 second audio, dividing file...")
            isSilence = True
        if isSilence:
            previousAudio = b''
            transcription.append("")
        transcription[segment] = data
        return transcription, segment, previousAudio, data, isSilence, seconds


    async def handleFile(self, data, context, record):
        text = self.preprocessRequest(data, record)
        result = self.transcribe(text)
        return result
    

    def handleRecord(self, data, transcription, previousAudio, segment, seconds, record):
        text, isSilence, previousAudio, seconds = self.preprocessStreaming(data, previousAudio, seconds, record)
        result = ""
        if not isSilence:
            result = self.transcribe(text, transcription[segment])
        transcription, segment, previousAudio, result, isSilence, seconds = self.postprocess(transcription, isSilence, segment, previousAudio, result, seconds)
        return result, transcription, previousAudio, segment, isSilence, seconds