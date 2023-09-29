package controllers

import (
	"context"
	"database/sql"
	usecase "simple-invitation/domain/usecase"
	"simple-invitation/protocol/grpc/presenters/compiled"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type memberController struct {
	memberUsecase usecase.IMemberUsecase
	compiled.UnimplementedMemberServiceServer
}

type IMemberController interface {
	FetchAll(context.Context, *compiled.MemberRequest) (*compiled.MemberResponse, error)
}

func NewMemberController(dbReader, dbWriter *sql.DB, grpcServer *grpc.Server) *memberController {
	var memberUsecase usecase.IMemberUsecase = usecase.NewMemberUsecase(dbReader, dbWriter)
	compiled.RegisterMemberServiceServer(grpcServer, &memberController{})
	reflection.Register(grpcServer)
	return &memberController{
		memberUsecase: memberUsecase,
	}
}

func (mc *memberController) FetchAll(context.Context, *compiled.MemberRequest) (*compiled.MemberResponse, error) {
	// data, err := mc.memberUsecase.FetchAll()

	res := &compiled.Member{
		Id:        "1",
		Firstname: "Shamir",
		Lastname:  "Husein",
		Email:     "shamirhusein@gmail.com",
	}
	data := compiled.MemberResponse{} //&compiled.MemberResponse{Members: data}
	data.Member = append(data.Member, res)
	return &data, nil

}
