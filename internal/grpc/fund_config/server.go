package fund_config

import (
	"context"
	"errors"
	"finance_service/internal/services/fund_config"
	"time"

	finv1 "github.com/aolychkin/protos/gen/go/finance"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Структура, реализующая функционал API
type serverAPI struct {
	finv1.UnimplementedFundConfigServer
	fund_config FundConfig
}

// Тот самый интерфейс, котрый мы передавали в grpcApp
// Каркасы для RPC-методов, которые мы будем использовать
type FundConfig interface {
	CreateFund(
		ctx context.Context,
		unionID int64,
		label string,
		priority uint64,
		check_child bool,
		rule_value int64,
		is_tmp bool,
	) (fund_id int64, err error)
	AddChild(
		ctx context.Context,
		fundID int,
		childID int,
	) (fund_id int, child_id int, err error)
	AddGoal(
		ctx context.Context,
		label string,
		total int64,
		expire_date time.Time,
		fundID int64,
		is_tmp bool,
	) (goal_id int, err error)
	GetFund(
		ctx context.Context,
		fundID int64,
	) (label string, auto_operation_id []int, child_id []int, goal_id []int, err error)
}

// Регистрируем serverAPI в gRPC-сервере
func Register(gRPC *grpc.Server, fund_config FundConfig) {
	finv1.RegisterFundConfigServer(gRPC, &serverAPI{fund_config: fund_config})
}

// SERVISE LAYOUT handler (интерфейс будущего Finance из сервисного слоя )
func (s *serverAPI) CreateFund(ctx context.Context, req *finv1.CreateFundRequest) (*finv1.CreateFundResponse, error) {
	if err := validateCreateFund(req); err != nil {
		return nil, err
	}

	fund_id, err := s.fund_config.CreateFund(ctx, req.GetUnionId(), req.GetLabel(), req.GetPriority(), req.GetCheckChild(), req.GetRuleValue(), req.GetIsTmp())

	if err != nil {
		if errors.Is(err, fund_config.ErrInvalidUnionID) {
			return nil, status.Error(codes.InvalidArgument, "invalid email or password")
		}

		return nil, status.Error(codes.Internal, "failed to login")
	}

	return &finv1.CreateFundResponse{
		FundId: fund_id,
	}, nil
}

// Validation functions
func validateCreateFund(req *finv1.CreateFundRequest) error {
	if req.GetUnionId() == 0 {
		return status.Error(codes.InvalidArgument, "union is required")
	}

	if req.GetLabel() == "" {
		return status.Error(codes.InvalidArgument, "fund label is required")
	}

	return nil
}
