package ai

import (
	"fmt"
	"go-xox-grpc-ai/internal/game"
	"math/rand"
	"time"
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
		aiFunction = aiRandomPlay(game)
	} else if difficulty == "2" {
		aiFunction = aiBestOrRandom(game)
	} else if difficulty == "3" {
		aiFunction = aiPlay(game)
	}

	for !game.CheckGameFinished() {
		if game.GetCurrentPlayer() == player {
			game.PlayRound()
		} else {
			aiFunction()
		}
		game.RenderBoard()
	}

	winner := game.GetWinner()
	if winner == "-" {
		fmt.Println("Tie!")
	} else {
		fmt.Println("Player " + game.GetWinner() + " Won!")
	}

}

func aiBestOrRandom(g *game.GameBoard) func() {
	return func() {
		rand.Seed(time.Now().UnixNano())
		random := rand.Intn(2)
		if random == 0 {
			aiRandomPlay(g)()
		} else {
			aiPlay(g)()
		}
	}
}

func aiRandomPlay(g *game.GameBoard) func() {
	return func() {
		var position int = -1
		for position == -1 {
			rand.Seed(time.Now().UnixNano())
			randomPos := rand.Intn(9)
			if g.GetBoard()[randomPos] == "" {
				position = randomPos
			}
		}

		g.SetBoardValue(position, g.GetCurrentPlayer())
		g.SwitchCurrentPlayer()
	}
}

func aiPlay(g *game.GameBoard) func() {
	return func() {
		pos := findBestMove(g.GetBoard(), g.GetCurrentPlayer())

		g.SetBoardValue(pos, g.GetCurrentPlayer())
		g.SwitchCurrentPlayer()
	}
}

func findBestMove(board []string, currentPlayer string) int {
	bestVal := -1000
	bestMove := -1

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			pos := i*3 + j%3
			if board[pos] == "" {
				board[pos] = currentPlayer

				moveValue := miniMax(board, 0, false, currentPlayer)

				board[pos] = ""

				if moveValue > bestVal {
					bestMove = pos
					bestVal = moveValue
				}
			}
		}
	}

	return bestMove
}

func isMovesLeft(board []string) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			pos := i*3 + j%3
			if board[pos] == "" {
				return true
			}
		}
	}
	return false
}

func miniMax(board []string, depth int, isMax bool, currentPlayer string) int {
	score := heuristic(board, currentPlayer) // evaluate heuristic result

	if score == 10 || score == -10 {
		return score
	}

	if !isMovesLeft(board) {
		return 0
	}

	if isMax {
		best := -1000

		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				pos := i*3 + j%3
				if board[pos] == "" {
					board[pos] = currentPlayer
					best = max(best, miniMax(board, depth+1, !isMax, currentPlayer))
					board[pos] = ""
				}
			}
		}
		return best
	} else {
		best := 1000

		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				pos := i*3 + j%3
				if board[pos] == "" {
					if currentPlayer == "O" {
						board[pos] = "X"
					} else {
						board[pos] = "O"
					}
					best = min(best, miniMax(board, depth+1, !isMax, currentPlayer))
					board[pos] = ""
				}
			}
		}
		return best
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func heuristic(board []string, currentPlayer string) int {
	opponent := "X"
	if currentPlayer == "X" {
		opponent = "O"
	}

	for i := 0; i < 3; i++ {
		if board[i*3] == board[i*3+1] && board[i*3+1] == board[i*3+2] {
			if board[i*3] == currentPlayer {
				return 10
			} else if board[i*3] == opponent {
				return -10
			}
		}
	}

	for i := 0; i < 3; i++ {
		if board[i] == board[3+i] && board[3+i] == board[6+i] {
			if board[i] == currentPlayer {
				return 10
			} else if board[i] == opponent {
				return -10
			}
		}
	}

	if board[0] == board[4] && board[4] == board[8] {
		if board[0] == currentPlayer {
			return 10
		} else if board[0] == opponent {
			return -10
		}
	}

	if board[2] == board[4] && board[4] == board[6] {
		if board[2] == currentPlayer {
			return 10
		} else if board[2] == opponent {
			return -10
		}
	}

	return 0
}
