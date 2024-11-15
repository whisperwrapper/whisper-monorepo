from grpcClient import GrpcClient

import random
import asyncio
import grpc
import time
import logging

class ConsolePrinter:

    def __init__(
            self, 
            host:str = None, 
            port:str = None, 
            language:str = None, 
            model:str = 'small',
            save:str = None
        ):
        self.grpcClient = GrpcClient(host, port, language, model, save)

    
    def _errorHandler(func):
        async def wrapper(*args, **kwargs):
            start = time.time()
            try:
                return await func(*args, **kwargs)
            except grpc.RpcError as grpcError:
                logging.error(f'Grpc connection failure: {grpcError.details()}') #TODO: when there is implemented secure connection, here handle those exceptions
            except Exception as e:
                logging.error(f'This is an unhandled exception: {e}')
            finally:
                end = time.time()
                logging.info(f'Execution time: {end - start}')
        return wrapper
    

    @_errorHandler
    async def startApp(self):
        seed = str(random.randint(0, 10000))
        conTask = asyncio.create_task(self.grpcClient.initiateConnection(seed)) # Initiate connection attempt async
        dot = 0
        while not conTask.done(): # Dot animation until connection task is finished
            dot = self._waitingAnimation(dot)
            await asyncio.sleep(0.1)
        self._waitingAnimation(3)
        await conTask
        connected = conTask.result()
        if connected.text != seed: # Unsuccessful if server returned different number
            logging.error("Problem connecting to the server")
            return False
        
        logging.info("Connected to the server!")
        return True
    

    @_errorHandler
    async def sendFile(self, audio:bytes):
        sendTask = asyncio.create_task(self.grpcClient.sendSoundFile(audio)) # Initiate sending file async
        dot = 0
        while not sendTask.done(): # Dot animation until connection task is finished
            dot = self._waitingAnimation(dot)
            await asyncio.sleep(0.1)
        self._waitingAnimation(3)
        script = sendTask.result() # Received transcribed text 
        print(script)
    

    @_errorHandler
    async def record(self):
        await self.grpcClient.streamSoundFile() # Initiate streaming file async


    def _waitingAnimation(self, dot: int) -> int:
        if dot == 3:
            print(end='\b\b\b', flush=True)
            print(end='   ',    flush=True)
            print(end='\b\b\b', flush=True)
            dot = 0
        else:
            print(end='.', flush=True)
            dot += 1
        return dot