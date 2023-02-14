package online

import (
	"context"
	"fmt"
	"go-xox-grpc-ai/internal/game"
	"go-xox-grpc-ai/internal/game/online/api"
	"io"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	grpcClient   api.GameServiceClient
	currentGame  *game.GameBoard
	clientPlayer string
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) JoinGame(host string) {
	conn, err := grpc.Dial(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	c.grpcClient = api.NewGameServiceClient(conn)

	resp, err := c.grpcClient.Join(context.Background(), &api.JoinRequest{})
	if err != nil || (resp != nil && !resp.Success) {
		fmt.Println("Host did not accepted your join request")
		os.Exit(1)
	}

	fmt.Println("\nYou have joined the game!")
	c.currentGame = game.NewGame()
	c.currentGame.Start()
	fmt.Printf("\nWaiting your opponent for first move...\n\n")
	fmt.Printf("==========================\n")

	c.ReadStream()
}

func (c *Client) ReadStream() {
	ctx := context.Background()

	waitResponse := make(chan error)

	stream, _ := c.grpcClient.ServerMove(ctx, &api.ServerMoveRequest{})
	c.clientPlayer = "O"

	for {
		streamResponse, err := stream.Recv()
		if err == io.EOF {
			waitResponse <- nil
			return
		}
		if err != nil {
			waitResponse <- fmt.Errorf("cannot receive stream response: %v", err)
			return
		}

		c.MovePlayed(streamResponse.Board, streamResponse.IsGameFinished)

		for c.currentGame.GetCurrentPlayer() == c.clientPlayer && !c.currentGame.IsFinished() {
			choiceAsInt, err := MovePosition()
			if err != nil {
				fmt.Println("Please select a legal position")
				continue
			}

			req := &api.ClientMoveRequest{
				Position: int32(choiceAsInt),
			}

			resp, err := c.grpcClient.ClientMove(context.Background(), req)
			if err != nil {
				fmt.Println("Error:", err)
			}
			if resp.Success {
				c.MovePlayed(resp.Board, streamResponse.IsGameFinished)
			} else {
				fmt.Println("Please select a legal position")
			}
		}
	}
}

func (c *Client) MovePlayed(board []string, isFinished bool) {
	c.currentGame.SwitchCurrentPlayer()
	c.currentGame.SetBoard(board)
	c.currentGame.RenderBoard()

	if isFinished {
		c.currentGame.CheckGameFinished()
		winner := c.currentGame.GetWinner()
		if winner == "-" {
			fmt.Println("Tie!")
		} else {
			fmt.Println("Player " + winner + " Won!")
		}
	}
}
