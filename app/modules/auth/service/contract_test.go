package service

import (
	"context"
	"testing"

	"github.com/adamnasrudin03/go-template/app/configs"
	"github.com/adamnasrudin03/go-template/app/modules/auth/repository/mocks"
	mocksLog "github.com/adamnasrudin03/go-template/app/modules/log/repository/mocks"
	mocksUser "github.com/adamnasrudin03/go-template/app/modules/user/repository/mocks"
	"github.com/adamnasrudin03/go-template/pkg/driver"
	"github.com/stretchr/testify/suite"
)

type AuthServiceTestSuite struct {
	suite.Suite
	repo     *mocks.AuthRepository
	repoUser *mocksUser.UserRepository
	repoLog  *mocksLog.LogRepository
	ctx      context.Context
	service  AuthService
}

func (srv *AuthServiceTestSuite) SetupTest() {
	var (
		cfg    = configs.GetInstance()
		logger = driver.Logger(cfg)
	)
	srv.ctx = context.Background()

	srv.repo = &mocks.AuthRepository{}
	srv.repoUser = &mocksUser.UserRepository{}
	srv.repoLog = &mocksLog.LogRepository{}
	params := AuthSrv{
		Repo:     srv.repo,
		RepoUser: srv.repoUser,
		RepoLog:  srv.repoLog,
		Cfg:      cfg,
		Logger:   logger,
	}

	srv.service = NewAuthService(params)
}

func TestAuthService(t *testing.T) {
	suite.Run(t, new(AuthServiceTestSuite))
}
