import grpc

import sakuraTalk_pb2_grpc, sakuraTalk_pb2

with grpc.insecure_channel('localhost:8806') as channel:
    stub = sakuraTalk_pb2_grpc.TalkServiceStub(channel)
    response = stub.deleteFriends(sakuraTalk_pb2.deleteFriendsRequest(mid="fsdf"))

