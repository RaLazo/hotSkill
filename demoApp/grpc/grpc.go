package grpc

import (
	"log"
	"net"

	"github.com/RaLazo/hotSkill/natsApp/grpc/chat"
	"google.golang.org/grpc"
)

type Runner struct{}

func New() Runner {
	return Runner{}
}

func (r Runner) Run() {
	log.Println("Starting Runner: GRPC Server localhost:9000")
	// create a TCP Listner on Port 9000
	lis, err := net.Listen("tcp", ":9000")
	// this how you handle errors in Golang
	if err != nil {
		log.Fatalf("Failed to listen on port 9000 %v", err)
	}

	// this is just a structure that has an interface with needed function SayHello
	// have a look at /chat/chat.go for more information
	s := chat.Server{}

	//create the GRCP Server
	grpcServer := grpc.NewServer()

	//
	chat.RegisterGameServiceServer(grpcServer, &s)

	// start listening on port 9000 for rpc
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000 %v", err)
	}
}
