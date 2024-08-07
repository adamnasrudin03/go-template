package service

import (
	"context"

	response_mapper "github.com/adamnasrudin03/go-helpers/response-mapper/v1"

	"github.com/adamnasrudin03/go-template/app/configs"
	"github.com/adamnasrudin03/go-template/app/modules/log/dto"
	"github.com/adamnasrudin03/go-template/app/modules/log/repository"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type LogService interface {
	GetList(ctx context.Context, params *dto.ListLogReq) (*response_mapper.Pagination, error)
	CreateByMessage(ctx context.Context, message string) (err error)
	Download(ctx *gin.Context, params *dto.ListLogReq) (err error)
}

type logSrv struct {
	Repo   repository.LogRepository
	Cfg    *configs.Configs
	Logger *logrus.Logger
}

func NewLogService(
	logRepo repository.LogRepository,
	cfg *configs.Configs,
	logger *logrus.Logger,
) LogService {
	return &logSrv{
		Repo:   logRepo,
		Cfg:    cfg,
		Logger: logger,
	}
}
