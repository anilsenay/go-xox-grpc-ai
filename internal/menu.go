package internal

import (
	"fmt"
	"go-xox-grpc-ai/internal/game/ai"
)

func MainMenu() {
	fmt.Println("Welcome to Tic Tac Toe!")
	fmt.Println("Please select your choice:")
	fmt.Println("1) Play against computer")
	fmt.Println("2) Play against your friend")
	var choice string
	for choice != "1" && choice != "2" {
		fmt.Print("Choice: ")
		choice = getUserInput()
	}

	if choice == "1" {
		playAgainstComputer()
	} else if choice == "2" {
		playAgainstPlayer()
	}
}

func playAgainstComputer() {
	fmt.Println()
	fmt.Println("1) Easy")
	fmt.Println("2) Hard")
	var choice string
	for choice != "1" && choice != "2" {
		fmt.Print("Difficulty: ")
		choice = getUserInput()
	}
	ai.StartGame(choice)
}

func playAgainstPlayer() {
	fmt.Println()
	fmt.Println("1) Host a game")
	fmt.Println("2) Join a game")
	var choice string
	for choice != "1" && choice != "2" {
		fmt.Print("Choice: ")
		choice = getUserInput()
	}

	if choice == "1" {
		// host game
	} else if choice == "2" {
		// join game
	}
}

func getUserInput() string {
	var choice string
	fmt.Scanln(&choice)
	return choice
}
