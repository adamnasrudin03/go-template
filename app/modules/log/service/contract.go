package service

import (
	"context"

	"github.com/adamnasrudin03/go-template/app/modules/log/payload"
	"github.com/adamnasrudin03/go-template/app/modules/log/repository"
	"github.com/adamnasrudin03/go-template/pkg/helpers"
	"github.com/sirupsen/logrus"
)

type LogService interface {
	GetList(ctx context.Context, params *payload.ListLogReq) (*helpers.Pagination, error)
	CreateByMessage(ctx context.Context, message string) (err error)
}

type logSrv struct {
	Repo   repository.LogRepository
	Logger *logrus.Logger
}

func NewLogService(
	logRepo repository.LogRepository,
	logger *logrus.Logger,
) LogService {
	return &logSrv{
		Repo:   logRepo,
		Logger: logger,
	}
}
