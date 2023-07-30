package main

import (
	"fmt"
	"time"

	pb "github.com/akifkadioglu/protocol-buffers-golang/proto"
)

func doDummy() *pb.Dummy {
	return &pb.Dummy{
		Id: 5,
	}
}

func main() {
	fmt.Println(doDummy())
	time.Sleep(time.Second * 10)
}
