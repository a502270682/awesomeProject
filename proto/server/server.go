package main

import (
	"awesomeProject/proto"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":6001")
	if err != nil {
		return
	}
	s := grpc.NewServer()
	proto.RegisterPeopleServiceServer(s, proto.CreateStudent{})
	fmt.Println("server running...")
	s.Serve(lis)
}


