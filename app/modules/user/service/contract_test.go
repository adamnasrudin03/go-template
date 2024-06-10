package service

import (
	"context"
	"testing"

	"github.com/adamnasrudin03/go-template/app/configs"
	"github.com/adamnasrudin03/go-template/app/modules/user/repository/mocks"
	"github.com/adamnasrudin03/go-template/pkg/driver"
	"github.com/stretchr/testify/suite"
)

type UserServiceTestSuite struct {
	suite.Suite
	repo    *mocks.UserRepository
	ctx     context.Context
	service UserService
}

func (srv *UserServiceTestSuite) SetupTest() {
	var (
		cfg    = configs.GetInstance()
		logger = driver.Logger(cfg)
	)

	srv.repo = &mocks.UserRepository{}
	srv.ctx = context.Background()
	srv.service = NewUserService(
		srv.repo,
		cfg,
		logger,
	)
}

func TestUserService(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}
