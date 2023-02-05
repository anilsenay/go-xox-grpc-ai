package online

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/peer"

	"go-xox-grpc-ai/internal/game/online/api"
	"go-xox-grpc-ai/internal/utils"
)

type gameService struct {
	api.UnimplementedGameServiceServer
}

func (s *gameService) Join(ctx context.Context, in *api.JoinRequest) (*api.JoinResponse, error) {
	p, _ := peer.FromContext(ctx)

	fmt.Printf("A player with ip %s want to join. Do you accept? (Y/N): ", p.Addr.String())
	var choice = utils.GetUserInput("Y", "N")

	var success bool
	if choice == "Y" {
		success = true
	}

	return &api.JoinResponse{
		Success: success,
	}, nil
}

func HostGame() {
	lis, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	api.RegisterGameServiceServer(s, &gameService{})

	fmt.Printf("server hosting at %v\n", lis.Addr().String())
	fmt.Println("waiting opponent...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// Client

func JoinGame(host string) {
	conn, err := grpc.Dial(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	client := api.NewGameServiceClient(conn)

	resp, err := client.Join(context.Background(), &api.JoinRequest{})
	if err != nil || (resp != nil && !resp.Success) {
		fmt.Println("Host did not accepted your join request")
		os.Exit(1)
	}

	fmt.Println("You have joined the game!")
	fmt.Println("Waiting your opponent to start the game...")
}
