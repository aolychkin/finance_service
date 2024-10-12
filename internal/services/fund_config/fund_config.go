package fund_config

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"finance_service/internal/domain/models"
	"finance_service/internal/lib/logger/sl"
	"finance_service/internal/storage"
)

type FundConfig struct {
	log          *slog.Logger
	fcfgSaver    FundConfigSaver
	fcfgProvider FundConfigProvider
}

// AddChild implements fund_config.FundConfig.
func (a *FundConfig) AddChild(ctx context.Context, fundID int, childID int) (fund_id int, child_id int, err error) {
	panic("unimplemented")
}

// AddGoal implements fund_config.FundConfig.
func (a *FundConfig) AddGoal(ctx context.Context, label string, total int64, expire_date time.Time, fundID int64, is_tmp bool) (goal_id int, err error) {
	panic("unimplemented")
}

// GetFund implements fund_config.FundConfig.
func (a *FundConfig) GetFund(ctx context.Context, fundID int64) (label string, auto_operation_id []int, child_id []int, goal_id []int, err error) {
	panic("unimplemented")
}

// Интерфейс хранилища
type FundConfigSaver interface {
	SaveFundConfig(
		ctx context.Context,
		unionID int64,
		label string,
		priority uint64,
		check_child bool,
		rule_value int64,
		is_tmp bool,
	) (fund_id int64, err error)
}

// Интерфейс хранилища
type FundConfigProvider interface {
	FundConfig(
		ctx context.Context,
		fundID int64,
	) (
		models.Fund,
		error,
	)
}

var (
	ErrInvalidUnionID = errors.New("invalid union id")
)

// Конструктор для нашего сервиса FundConfig
func New(
	log *slog.Logger,
	fundConfigSaver FundConfigSaver,
	fundConfigProvider FundConfigProvider,
) *FundConfig {
	return &FundConfig{
		fcfgSaver:    fundConfigSaver,
		fcfgProvider: fundConfigProvider,
		log:          log,
	}
}

// TODO: И теперь ошибки можно логировать вот таким образом:
// Login checks if user with given credentials exists in the system.
func (a *FundConfig) CreateFund(
	ctx context.Context,
	unionID int64,
	label string,
	priority uint64,
	check_child bool,
	rule_value int64,
	is_tmp bool,
) (int64, error) {
	const op = "FundConfig.CreateFund"

	log := a.log.With(
		slog.String("op", op),
		slog.String("FundLabel", label),
	)

	log.Info("adding new fund config")

	fund_id, err := a.fcfgSaver.SaveFundConfig(ctx, unionID, label, priority, check_child, rule_value, is_tmp)
	if err != nil {
		if errors.Is(err, storage.ErrInvalidUnionID) {
			log.Warn("user already exists", sl.Err(err))

			return 0, fmt.Errorf("%s: %w", op, ErrInvalidUnionID)
		}

		log.Error("failed to save user", sl.Err(err))

		return 0, fmt.Errorf("%s: %w", op, err)
	}

	log.Info("fund config added successfully")

	return fund_id, nil
}
