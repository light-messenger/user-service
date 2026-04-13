package handler

import (
	"context"

	pb "github.com/light-messenger/user-service/pkg/userservice"
)

func (h *Handler) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	errValidate := validateCreateRequest(req)
	if errValidate != nil {
		return nil, errValidate
	}

	id, err := h.service.Create(ctx, req.GetNickname(), req.GetPassword())
	if err != nil {
		return nil, err
	}

	return &pb.CreateResponse{Id: id}, nil
}

func validateCreateRequest(req *pb.CreateRequest) error {
	if req.GetNickname() == "" {
		return errEmptyNickname
	}

	if req.GetPassword() == "" {
		return errEmptyPassword
	}

	return nil
}
