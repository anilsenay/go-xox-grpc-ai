package ai

import (
	"fmt"
	"go-xox-grpc-ai/internal/game"
)

var Difficulty string

func StartGame(difficulty string) {
	Difficulty = difficulty

	fmt.Println("Want to be X or O? ")
	var player string
	for player != "X" && player != "O" {
		fmt.Print("Choice: ")
		fmt.Scanln(&player)
	}

	game := game.NewGame()
	game.Start()

	var aiFunction func()

	if difficulty == "1" {
		aiFunction = game.AiRandomPlay
	} else if difficulty == "2" {
		// implement minmax algorithm later
		aiFunction = game.AiRandomPlay
	}

	for !game.CheckGameFinished() {
		if game.GetCurrentPlayer() == player {
			game.PlayRound()
		} else {
			aiFunction()
		}
		game.RenderBoard()
	}

	fmt.Println("Player " + game.GetWinner() + " Won!")

}
