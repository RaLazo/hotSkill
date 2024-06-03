package chat

import (
	"log"

	"github.com/RaLazo/hotSkill/natsApp/game"
	"golang.org/x/net/context"
)

type Server struct{}

// mustEmbedUnimplementedGameServiceServer implements GameServiceServer.
func (s *Server) mustEmbedUnimplementedGameServiceServer() {
	panic("unimplemented")
}

func (s *Server) VerifyWinner(ctx context.Context, message *Message) (*Message, error) {
	log.Println("CheckIn for Winner")
	return &Message{Body: game.VerfiyWinner(message.Body)}, nil
}
