package talkServer

import (
	"github.com/sakura-rip/SakuraTalk/TalkService"
	"google.golang.org/grpc"
	"net"
)

type TalkHandler struct{}

// RunServer サーバーを走らせます！
func RunServer() {
	listen, err := net.Listen("tcp", ":8806")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	talkService := TalkHandler{}
	TalkService.RegisterTalkServiceServer(server, talkService)
	if err := server.Serve(listen); err != nil {
		panic(err)
	}
}
