package service

import (
	"context"

	"github.com/adamnasrudin03/go-template/app/modules/log/payload"
	"github.com/adamnasrudin03/go-template/app/modules/log/repository"
	"github.com/adamnasrudin03/go-template/pkg/helpers"
)

type LogService interface {
	GetList(ctx context.Context, params *payload.ListLogReq) (*helpers.Pagination, error)
}

type logSrv struct {
	Repo repository.LogRepository
}

func NewLogService(
	logRepo repository.LogRepository,
) LogService {
	return &logSrv{
		Repo: logRepo,
	}
}
