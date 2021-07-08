import os
import uuid

def uuidV1():
    #return uuid.uuid1(bytes=os.urandom(16))
    return uuid.UUID(bytes=os.urandom(16), version=1)