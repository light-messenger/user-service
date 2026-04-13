package handler

import (
	"context"

	pb "github.com/light-messenger/user-service/pkg/userservice"
)

func (h *Handler) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	errValidate := validateGetRequest(req)
	if errValidate != nil {
		return nil, errValidate
	}

	nickname, err := h.service.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.GetResponse{Nickname: nickname}, nil
}

func validateGetRequest(req *pb.GetRequest) error {
	if req.GetId() <= 0 {
		return errIncorrectId
	}

	return nil
}
