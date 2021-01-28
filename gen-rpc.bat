cd /d %~dp0
protoc --go_out=plugins=grpc:./server/TalkService sakuraTalk.proto
git add ./server/TalkService
git commit -m "rpc: generated commit"

python -m grpc.tools.protoc -I. --python_out=./client/TalkService --grpc_python_out=./client/TalkService sakuraTalk.proto