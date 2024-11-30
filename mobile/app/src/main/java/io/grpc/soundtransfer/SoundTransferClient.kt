package io.grpc.soundtransfer

import android.net.Uri
import android.util.Log
import com.google.protobuf.kotlin.toByteString
import io.grpc.ManagedChannelBuilder
import io.grpc.Metadata
import jWT
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.asExecutor
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.launch
import java.io.Closeable
import java.io.File

class SoundTransferClient(uri: Uri) : Closeable {
    private val audiStreamManager: AudioStreamManager = AudioStreamManager()
    private val channel = let {
        println("Connecting to ${uri.host}:${uri.port}")

        val builder = ManagedChannelBuilder.forAddress(uri.host, uri.port)
        if (uri.scheme == "https") {
            builder.useTransportSecurity()
        } else {
            builder.usePlaintext()
        }

        builder.executor(Dispatchers.IO.asExecutor()).build()
    }
    private val transferer = SoundServiceGrpcKt.SoundServiceCoroutineStub(channel)
    val stub = SoundServiceGrpcKt.SoundServiceCoroutineStub(channel)

    suspend fun sendSoundFile(filePath: String): String? {
        try {
            val metadata = Metadata()
            val key = Metadata.Key.of("JWT", Metadata.ASCII_STRING_MARSHALLER)
            if (jWT != "") {
                metadata.put(key, jWT)
            }
            val bytes = File(filePath).readBytes().toByteString()
            val request = soundRequest { this.soundData = bytes }
            val response = transferer.sendSoundFile(request, metadata)
            return response.text
        } catch (e: Exception) {
            e.printStackTrace()
        }
        return null
    }

    fun streamSoundFile() {
        val metadata = Metadata()
        val key = Metadata.Key.of("language", Metadata.ASCII_STRING_MARSHALLER)
        metadata.put(key, "pl")
        audiStreamManager.initAudioRecorder()
        audiStreamManager.record()
        CoroutineScope(Dispatchers.Default).launch {
            val requests: Flow<SoundRequest> = audiStreamManager.record()
            stub.streamSoundFile(requests, metadata).collect { response ->
                Log.i("stream", "Got message: \"${response.text}\"")
            }
        }
    }

    fun stopStream() {
        audiStreamManager.stop()
    }

    override fun close() {
        channel.shutdownNow()
    }

}