package main

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/akifkadioglu/protocol-buffers-golang/greet/proto"
)

type Server struct {
	pb.GreetServiceServer
}

const (
	addr string = "0.0.0.0:8081"
)

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer conn.Close()
	c := pb.NewGreetServiceClient(conn)
	doGreet(c)
}
func doGreet(c pb.GreetServiceClient) {
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		Id: uuid.New().String(),
	})
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Printf(res.Text)
}
