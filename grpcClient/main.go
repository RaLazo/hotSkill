package main

import (
	"log"

	"golang.org/x/net/context"

	"github.com/RaLazo/hotSkill/grpcClient/chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct{}

func main() {
	// create client for GRPC Server
	var conn *grpc.ClientConn
	conn, err := grpc.NewClient(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("no server connection could be established cause: %v", err)
	}

	// defer runs after the functions finishes
	defer conn.Close()

	c := chat.NewGameServiceClient(conn)

	message := chat.Message{
		Body: "S",
	}

	response, err := c.VerifyWinner(context.Background(), &message)
	log.Printf("Response from Server: %s\n", response.Body)
}
