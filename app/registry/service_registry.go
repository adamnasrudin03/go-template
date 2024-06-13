package registry

import (
	"github.com/adamnasrudin03/go-template/app/configs"
	authSrv "github.com/adamnasrudin03/go-template/app/modules/auth/service"
	logSrv "github.com/adamnasrudin03/go-template/app/modules/log/service"
	userSrv "github.com/adamnasrudin03/go-template/app/modules/user/service"
	"github.com/adamnasrudin03/go-template/pkg/driver"
	"github.com/sirupsen/logrus"
)

// Services all service object injected here
type Services struct {
	Auth authSrv.AuthService
	User userSrv.UserService
	Log  logSrv.LogService
}

func WiringService(repo *Repositories, cache *driver.RedisClient, cfg *configs.Configs, logger *logrus.Logger) *Services {

	return &Services{
		Auth: regAuthSrv(repo, cfg, logger),
		User: userSrv.NewUserService(repo.User, cfg, logger),
		Log:  logSrv.NewLogService(repo.Log, cfg, logger),
	}
}

func regAuthSrv(repo *Repositories, cfg *configs.Configs, logger *logrus.Logger) authSrv.AuthService {
	params := authSrv.AuthSrv{
		Repo:      repo.Auth,
		RepoCache: repo.Cache,
		RepoUser:  repo.User,
		RepoLog:   repo.Log,
		Cfg:       cfg,
		Logger:    logger,
	}
	return authSrv.NewAuthService(params)
}
