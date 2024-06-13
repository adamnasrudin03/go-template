package service

import (
	"context"
	"testing"

	"github.com/adamnasrudin03/go-template/app/configs"
	mocksCache "github.com/adamnasrudin03/go-template/app/modules/cache/repository/mocks"
	mocksLog "github.com/adamnasrudin03/go-template/app/modules/log/repository/mocks"
	"github.com/adamnasrudin03/go-template/app/modules/user/repository/mocks"
	"github.com/adamnasrudin03/go-template/pkg/driver"
	"github.com/stretchr/testify/suite"
)

type UserServiceTestSuite struct {
	suite.Suite
	repo      *mocks.UserRepository
	repoCache *mocksCache.CacheRepository
	repoLog   *mocksLog.LogRepository
	ctx       context.Context
	service   UserService
}

func (srv *UserServiceTestSuite) SetupTest() {
	var (
		cfg    = configs.GetInstance()
		logger = driver.Logger(cfg)
	)

	srv.repo = &mocks.UserRepository{}
	srv.repoCache = &mocksCache.CacheRepository{}
	srv.repoLog = &mocksLog.LogRepository{}
	srv.ctx = context.Background()
	params := UserSrv{
		Repo:      srv.repo,
		RepoCache: srv.repoCache,
		RepoLog:   srv.repoLog,
		Cfg:       cfg,
		Logger:    logger,
	}
	srv.service = NewUserService(params)
}

func TestUserService(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}
