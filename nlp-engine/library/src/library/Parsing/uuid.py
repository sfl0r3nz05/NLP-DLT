import os
import uuid

def uuidV1():
    """
    Method used generate unique identifiers.
    """
    return uuid.UUID(bytes=os.urandom(16), version=1)