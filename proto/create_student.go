package proto

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"runtime"
	"time"
)

type CreateStudent struct {

}

func (c CreateStudent) mustEmbedUnimplementedPeopleServiceServer() {
	fmt.Println("not implemented")
}

func (c CreateStudent) CreatePeople(ctx context.Context, in *CreatePeopleReq) (*CreatePeopleRsp, error) {
	data := make(chan *CreatePeopleRsp, 1)
	fmt.Println("get calling")
	go handle(ctx, in, data)
	fmt.Println("start to ret")
	select {
	case <- ctx.Done():
		fmt.Println("outer ctx quit")
		return nil, errors.New("ctx done")
	case res := <- data:
		return res, nil
	}
}

func handle(ctx context.Context, req *CreatePeopleReq, data chan<- *CreatePeopleRsp) {
	select {
	case <-ctx.Done():
		fmt.Println("ctx done", ctx.Err())
		runtime.Goexit() //超时后退出该Go协程
	case <-time.After(4 * time.Second): // 模拟耗时操作
		res := CreatePeopleRsp{
			Ok: 0,
			Message: fmt.Sprintf("create %s success, age:%d", req.Name, req.Age),
		}
		// //修改数据库前进行超时判断
		// if ctx.Err() == context.Canceled{
		// 	...
		// 	//如果已经超时，则退出
		// }
		data <- &res
	}
}