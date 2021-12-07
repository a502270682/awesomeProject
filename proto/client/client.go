package main

import (
	"awesomeProject/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:6001", grpc.WithInsecure())
	if err != nil {
		return
	}
	defer conn.Close()
	c := proto.NewPeopleServiceClient(conn)
	fmt.Println("start", time.Now())

	// 自带关闭连接请求，避免grpc server端长链接一直没有响应，拖跨client端
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
	_, err = c.CreatePeople(ctx, &proto.CreatePeopleReq{
		Name: "gfy",
		Age:  100,
	})
	if err != nil {
		//获取错误状态
		state, ok := status.FromError(err)
		if ok {
			//判断是否为调用超时
			if state.Code() == codes.DeadlineExceeded {
				fmt.Println("Route timeout!", time.Now())
				cancel()
			}
		}
		cancel()
		fmt.Println("Call Route err: ", err)
	}
	fmt.Println("end", time.Now())
}
