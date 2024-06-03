package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

func verfiyWinner(userAction string, computerAction string) {
	if userAction == computerAction {
		fmt.Println("Equal no winner")
	}
	if userAction == "S" {
		if computerAction == "P" {
			fmt.Println("Computer: Paper")
			fmt.Println("User Won")
		} else {
			fmt.Println("Computer: Rock")
			fmt.Println("Computer Won")
		}
		return
	}

	if userAction == "P" {
		if computerAction == "R" {
			fmt.Println("Computer: Rock")
			fmt.Println("User Won")
		} else {
			fmt.Println("Computer: Scissors")
			fmt.Println("Computer Won")
		}
		return
	}
	if userAction == "R" {
		if computerAction == "S" {
			fmt.Println("Computer: Scissors")
			fmt.Println("User Won")
		} else {
			fmt.Println("Computer: Paper")
			fmt.Println("Computer Won")
		}
		return
	}
}

func getComputerAction() string {
	reasons := []string{
		"R",
		"P",
		"S",
	}
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	i := r.Intn(len(reasons))
	return reasons[i]
}

func getUserActionCli() string {
	reader := bufio.NewReader(os.Stdin)
	cliInput, _, err := reader.ReadRune()
	if err != nil {
		log.Fatal(err)
	}
	return strings.ToUpper(string(cliInput))
}

func userAction(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
	req.URL.Query()
}

func headers(w http.ResponseWriter, req *http.Request) {
	req.URL.Query()
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}
func main() {
	http.HandleFunc("/user/action", userAction)
	http.HandleFunc("/headers", headers)
	http.ListenAndServe(":8090", nil)
	// for {
	// 	fmt.Println("Rock-paper-scissors app")
	// 	fmt.Printf("User: ")
	// 	userAction := getUserActionCli()
	// 	computerAction := getComputerAction()
	// 	verfiyWinner(userAction, computerAction)
	// 	fmt.Println()
	// }
}
