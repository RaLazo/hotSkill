package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/RaLazo/hotSkill/natsApp/cli"
	"github.com/RaLazo/hotSkill/natsApp/grpc"
	"github.com/RaLazo/hotSkill/natsApp/http"
	"github.com/RaLazo/hotSkill/natsApp/runner"
)

func main() {
	done := make(chan os.Signal, 1)
	services := runner.MultiService{
		http.New(),
		cli.New(),
		grpc.New(),
	}

	log.Println("Starting Game Server")
	services.Start()

	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	log.Println("To quit, press ctrl+c")
	<-done
}
