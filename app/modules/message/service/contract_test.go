package service

import (
	"context"
	"testing"

	"github.com/adamnasrudin03/go-template/app/configs"
	mocksLog "github.com/adamnasrudin03/go-template/app/modules/log/repository/mocks"
	"github.com/adamnasrudin03/go-template/app/modules/message/repository/mocks"
	"github.com/adamnasrudin03/go-template/pkg/driver"
	"github.com/stretchr/testify/suite"
)

type MessageServiceTestSuite struct {
	suite.Suite
	repo    *mocks.MessageRepository
	repoLog *mocksLog.LogRepository
	ctx     context.Context
	service MessageService
}

func (srv *MessageServiceTestSuite) SetupTest() {
	var (
		cfg    = configs.GetInstance()
		logger = driver.Logger(cfg)
	)

	srv.repo = &mocks.MessageRepository{}
	srv.repoLog = &mocksLog.LogRepository{}
	srv.ctx = context.Background()
	params := MessageSrv{
		Repo:    srv.repo,
		RepoLog: srv.repoLog,
		Cfg:     cfg,
		Logger:  logger,
	}
	srv.service = NewMessageService(params)
}

func TestMessageService(t *testing.T) {
	suite.Run(t, new(MessageServiceTestSuite))
}
