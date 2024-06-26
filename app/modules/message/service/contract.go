package service

import (
	"context"

	"github.com/adamnasrudin03/go-template/app/configs"
	logRepo "github.com/adamnasrudin03/go-template/app/modules/log/repository"
	"github.com/adamnasrudin03/go-template/app/modules/message/repository"
	"github.com/sirupsen/logrus"
)

type MessageService interface {
	Consume(ctx context.Context)
	CreateLogByMessage(ctx context.Context, message string) (err error)
}

type MessageSrv struct {
	Repo    repository.MessageRepository
	RepoLog logRepo.LogRepository
	Cfg     *configs.Configs
	Logger  *logrus.Logger
}

func NewMessageService(params MessageSrv) MessageService {
	return &params
}
