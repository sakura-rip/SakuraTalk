cd /d %~dp0
protoc --go_out=plugins=grpc:./server/talkService sakuraTalk.proto

python -m grpc.tools.protoc -I. --python_out=./client/protoc --grpc_python_out=./client/protoc sakuraTalk.proto
git add ./server/talkService
git add ./client/protoc/sakuraTalk_pb2.py
git add ./client/protoc/sakuraTalk_pb2_grpc.py
git commit -m "rpc: generated commit"
