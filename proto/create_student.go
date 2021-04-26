package proto

import (
	"context"
	"fmt"
)

type CreateStudent struct {

}

func (c CreateStudent) mustEmbedUnimplementedPeopleServiceServer() {
	fmt.Println("not implemented")
}

func (c CreateStudent) CreatePeople(ctx context.Context, in *CreatePeopleReq) (*CreatePeopleRsp, error) {
	return &CreatePeopleRsp{
		Ok: 0,
		Message: fmt.Sprintf("create %s success, age:%d", in.Name, in.Age),
	}, nil
}
