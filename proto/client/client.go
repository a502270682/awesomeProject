package main

import (
	"awesomeProject/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func main(){
	conn, err := grpc.Dial("localhost:6001", grpc.WithInsecure())
	if err != nil {
		return
	}
	defer  conn.Close()
	c := proto.NewPeopleServiceClient(conn)
	r, err := c.CreatePeople(context.Background(), &proto.CreatePeopleReq{
		Name: "gfy",
		Age: 100,
	})
	if err != nil {
		return
	}
	fmt.Println("client", r.Message, r.Ok)
}
