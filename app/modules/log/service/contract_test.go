package service

import (
	"context"
	"testing"

	"github.com/adamnasrudin03/go-template/app/configs"
	"github.com/adamnasrudin03/go-template/app/modules/log/repository/mocks"
	"github.com/adamnasrudin03/go-template/pkg/driver"
	"github.com/stretchr/testify/suite"
)

type LogServiceTestSuite struct {
	suite.Suite
	repoLog *mocks.LogRepository
	ctx     context.Context
	service LogService
}

func (srv *LogServiceTestSuite) SetupTest() {
	var (
		cfg    = configs.GetInstance()
		logger = driver.Logger(cfg)
	)

	srv.repoLog = &mocks.LogRepository{}

	srv.service = NewLogService(srv.repoLog, logger)
	srv.ctx = context.Background()
}

func TestLogService(t *testing.T) {
	suite.Run(t, new(LogServiceTestSuite))
}
