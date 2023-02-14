package online

import (
	"context"
	"fmt"
	"go-xox-grpc-ai/internal/game"
	"go-xox-grpc-ai/internal/game/online/api"
	"go-xox-grpc-ai/internal/utils"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

// protoc -I=. --go_out=api --go_opt=paths=source_relative --go-grpc_out=api --go-grpc_opt=paths=source_relative game.proto

type Server struct {
	api.UnimplementedGameServiceServer
	currentGame *game.GameBoard
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Join(ctx context.Context, in *api.JoinRequest) (*api.JoinResponse, error) {
	p, _ := peer.FromContext(ctx)

	fmt.Printf("A player with ip %s want to join. Do you accept? (Y/N): ", p.Addr.String())
	var choice = utils.GetUserInput("Y", "N")

	var success bool
	if choice == "Y" {
		success = true
		s.currentGame = game.NewGame()
		s.currentGame.Start()
	}

	return &api.JoinResponse{
		Success: success,
	}, nil
}

func (s *Server) HostGame() {
	lis, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		panic(err)
	}

	serv := grpc.NewServer()
	api.RegisterGameServiceServer(serv, s)

	fmt.Printf("server hosting at %v\n", lis.Addr().String())
	fmt.Println("waiting opponent...")
	if err := serv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *Server) ClientMove(ctx context.Context, req *api.ClientMoveRequest) (*api.ClientMoveResponse, error) {
	pos := int(req.Position)
	isLegalMove := s.currentGame.IsLegalMove(pos)
	if isLegalMove {
		s.MovePlayed(pos-1, "O")

		return &api.ClientMoveResponse{
			Success:        isLegalMove,
			Board:          s.currentGame.GetBoard(),
			CurrentPlayer:  s.currentGame.GetCurrentPlayer(),
			IsGameFinished: s.currentGame.CheckGameFinished(),
		}, nil
	} else {
		return nil, fmt.Errorf("please move a legal move")
	}
}

func (s *Server) ServerMove(req *api.ServerMoveRequest, stream api.GameService_ServerMoveServer) error {
	for {
		if s.currentGame.GetCurrentPlayer() != "X" {
			continue
		}

		pos, err := MovePosition()
		for err != nil {
			pos, err = MovePosition()
		}

		s.MovePlayed(pos-1, "X")

		stream.Send(&api.ServerMoveResponse{
			Position:       int32(pos),
			Board:          s.currentGame.GetBoard(),
			CurrentPlayer:  s.currentGame.GetCurrentPlayer(),
			IsGameFinished: s.currentGame.CheckGameFinished(),
		})
	}
}

func MovePosition() (int, error) {
	fmt.Print("Select your move position:")
	var choice = utils.GetUserInput("1", "2", "3", "4", "5", "6", "7", "8", "9")
	return strconv.Atoi(choice)
}

func (s *Server) MovePlayed(pos int, player string) {
	s.currentGame.SetBoardValue(pos, player)
	s.currentGame.RenderBoard()
	s.currentGame.SwitchCurrentPlayer()

	isFinished := s.currentGame.CheckGameFinished()
	if isFinished {
		s.currentGame.CheckGameFinished()
		winner := s.currentGame.GetWinner()
		if winner == "-" {
			fmt.Println("Tie!")
		} else {
			fmt.Println("Player " + winner + " Won!")
		}
	}
}
