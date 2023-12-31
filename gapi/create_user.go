package gapi

import (
	"context"
	"time"

	db "github.com/fauzanfebrian/simplebank/db/sqlc"
	"github.com/fauzanfebrian/simplebank/pb"
	"github.com/fauzanfebrian/simplebank/util"
	"github.com/fauzanfebrian/simplebank/val"
	"github.com/fauzanfebrian/simplebank/worker"
	"github.com/hibiken/asynq"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	if violations := validateCreateUserRequest(req); violations != nil {
		return nil, invalidArgumentError(violations)
	}

	hashedPassword, err := util.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	arg := db.CreateUserTXParams{
		CreateUserParams: db.CreateUserParams{
			Username:       req.GetUsername(),
			FullName:       req.GetFullName(),
			Email:          req.GetEmail(),
			HashedPassword: hashedPassword,
			Role:           req.GetRole(),
		},
		AfterCreate: func(user db.User) error {
			return server.sendVerifyEmail(ctx, user)
		},
	}
	txResult, err := server.store.CreateUserTx(ctx, arg)
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			return nil, status.Errorf(codes.AlreadyExists, err.Error())
		}
		return nil, status.Errorf(codes.Internal, "failed to create user: %s", err)
	}

	res := &pb.CreateUserResponse{
		User: convertUser(txResult.User),
	}
	return res, nil
}

func (server *Server) sendVerifyEmail(ctx context.Context, user db.User) error {
	opts := []asynq.Option{
		asynq.MaxRetry(10),
		asynq.ProcessIn(10 * time.Second),
		asynq.Queue(worker.CriticalQueue),
	}
	taskPayload := &worker.PayloadSendVerifyEmail{
		Username: user.Username,
	}
	err := server.taskDistributor.DistributeTaskSendVerifyEmail(ctx, taskPayload, opts...)
	if err != nil {
		return status.Errorf(codes.Internal, "failed to distribute task: %s", err)
	}

	return nil
}

func validateCreateUserRequest(req *pb.CreateUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, fieldViolation("username", err))
	}

	if err := val.ValidateRole(req.GetRole()); err != nil {
		violations = append(violations, fieldViolation("role", err))
	}

	if err := val.ValidatePassword(req.GetPassword()); err != nil {
		violations = append(violations, fieldViolation("password", err))
	}

	if err := val.ValidateFullName(req.GetFullName()); err != nil {
		violations = append(violations, fieldViolation("full_name", err))
	}

	if err := val.ValidateEmail(req.GetEmail()); err != nil {
		violations = append(violations, fieldViolation("email", err))
	}

	return violations
}
