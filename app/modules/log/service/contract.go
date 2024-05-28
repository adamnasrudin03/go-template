package service

import (
	"github.com/adamnasrudin03/go-template/app/modules/log/repository"
)

type LogService interface {
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
