package app

import (
	"log/slog"

	grpcapp "finance_service/internal/app/grpc"
	"finance_service/internal/services/fund_config"
	"finance_service/internal/storage/sqlite"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
) *App {
	storage, err := sqlite.New(storagePath)
	if err != nil {
		panic(err)
	}

	fundConfigService := fund_config.New(log, storage, storage)

	grpcApp := grpcapp.New(log, fundConfigService, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}
}
