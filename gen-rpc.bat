cd /d %~dp0
protoc --go_out=plugins=grpc:./server/TalkService sakuraTalk.proto

python -m grpc.tools.protoc -I. --python_out=./client/TalkService --grpc_python_out=./client/TalkService sakuraTalk.proto
git add ./server/TalkService
git add ./client/TalkService
git commit -m "rpc: generated commit"
