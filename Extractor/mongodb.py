from pymongo import *
from pymongo import MongoClient
from pymongo import errors


def connect_database():

     

    MONGODB_HOST = 'localhost'
    MONGODB_PORT = '27017'
    MONGODB_TIMEOUT = 5000
    
    URI_CONNECTION = "mongodb://root:1234@localhost:27017/twitterdb"
    
    try:
        client = MongoClient(URI_CONNECTION, serverSelectionTimeoutMS=MONGODB_TIMEOUT)
        client.server_info()
        print 'OK -- Connected to MongoDB at server %s' % (MONGODB_HOST)
        client.close()
    except errors.ServerSelectionTimeoutError as error:
        
        print 'Error with MongoDB connection: %s' % error

    except errors.ConnectionFailure as error:
        print 'Could not connect to MongoDB: %s' % error



connect_database()
