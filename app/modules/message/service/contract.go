package service

import (
	"github.com/adamnasrudin03/go-template/app/configs"
	"github.com/adamnasrudin03/go-template/app/modules/message/repository"
	"github.com/sirupsen/logrus"
)

type MessageService interface {
}

type MessageSrv struct {
	Repo   repository.MessageRepository
	Cfg    *configs.Configs
	Logger *logrus.Logger
}

func NewMessageService(params MessageSrv) MessageService {
	return &params
}
