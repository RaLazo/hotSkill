package game

import (
	"log"
	"math/rand"
	"time"
)

const USER = "User"
const COMPUTER = "Computer"
const EQUAL = "Equal"

func VerfiyWinner(userAction string) string {
	computerAction := getComputerAction()

	if userAction == computerAction {
		log.Println("Equal no winner")
		return EQUAL
	}
	if userAction == "S" {
		if computerAction == "P" {
			log.Println("Computer: Paper")
			log.Println("User Won")
			return USER
		} else {
			log.Println("Computer: Rock")
			log.Println("Computer Won")
			return COMPUTER
		}
	}

	if userAction == "P" {
		if computerAction == "R" {
			log.Println("Computer: Rock")
			log.Println("User Won")
			return USER
		} else {
			log.Println("Computer: Scissors")
			log.Println("Computer Won")
			return COMPUTER
		}

	}
	if userAction == "R" {
		if computerAction == "S" {
			log.Println("Computer: Scissors")
			log.Println("User Won")
			return USER
		} else {
			log.Println("Computer: Paper")
			log.Println("Computer Won")
			return COMPUTER
		}
	}
	return ""
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
