package handler

import (
	"context"

	pb "github.com/light-messenger/user-service/pkg/userservice"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *Handler) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	errValidate := validateGetRequest(req)
	if errValidate != nil {
		return nil, errValidate
	}

	user, err := h.service.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.GetResponse{User: &pb.User{
		Id:           user.Id,
		Nickname:     user.Nickname,
		RegisteredAt: timestamppb.New(user.RegisteredAt),
	}}, nil
}

func validateGetRequest(req *pb.GetRequest) error {
	if req.GetId() <= 0 {
		return errIncorrectId
	}

	return nil
}
