package main

import (
	"fmt"
	"log"
	"flag"
	"io"
	"google.golang.org/grpc"
	"net"
	mathspb "github.com/zachmandeville/grpcchanfun/api/maths"
)

var (
	port = flag.Int("port", 10000, "the server port")
)

type mathsServer struct {
	mathspb.UnimplementedMathsServer
	squares map[int32]int32
}

func (s *mathsServer) Squares (stream mathspb.Maths_SquaresServer)  error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Received %v\n", in.Number)
		square := in.Number * in.Number
		resp := &mathspb.SquaresResponse{
			Number: in.Number,
			Square: square,
		}
		if err := stream.Send(resp); err != nil {
			return err
		}
	}
}

func (s *mathsServer) Cubes (stream mathspb.Maths_CubesServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Received to cube: %v\n", in.Number)
		cube := in.Number * in.Number * in.Number
		resp := &mathspb.CubesResponse{
			Number: in.Number,
			Cube: cube,
		}
		if err := stream.Send(resp); err != nil {
			return err
		}
	}
}

func newServer() *mathsServer {
	return &mathsServer{}
}

func main () {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%v", *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	mathspb.RegisterMathsServer(grpcServer, newServer())
	log.Printf("server started at localhost:%v\n", *port)
	grpcServer.Serve(lis)
}
