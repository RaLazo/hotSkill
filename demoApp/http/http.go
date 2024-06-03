package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/RaLazo/hotSkill/natsApp/game"
)

type Runner struct{}

func userAction(w http.ResponseWriter, req *http.Request) {
	urlValues := req.URL.Query()
	userAction := urlValues.Get("action")
	result := game.VerfiyWinner(userAction)
	fmt.Fprintf(w, "The game Result is: "+result)
}

func (r Runner) Run() {
	log.Println("Starting Runner: HTTP Server localhost:8080")
	http.HandleFunc("/user/action", userAction)
	http.ListenAndServe(":8090", nil)
}

func New() Runner {
	return Runner{}
}
