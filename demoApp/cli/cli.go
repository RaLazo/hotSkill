package cli

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/RaLazo/hotSkill/natsApp/game"
)

type Runner struct{}

func New() Runner {
	return Runner{}
}

func GetUserActionCli() string {
	for {
		log.Println("Provide CLI Parameter (e.g. S,P or R)")
		reader := bufio.NewReader(os.Stdin)
		cliInput, _, err := reader.ReadRune()
		if err != nil {
			log.Fatal(err)
		}
		game.VerfiyWinner(strings.ToUpper(string(cliInput)))
	}
}

func (r Runner) Run() {
	log.Println("Starting Runner: CLI")
	GetUserActionCli()
}
