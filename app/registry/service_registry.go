package registry

import (
	"github.com/adamnasrudin03/go-template/app/configs"
	authSrv "github.com/adamnasrudin03/go-template/app/modules/auth/service"
	logSrv "github.com/adamnasrudin03/go-template/app/modules/log/service"
	messageSrv "github.com/adamnasrudin03/go-template/app/modules/message/service"
	userSrv "github.com/adamnasrudin03/go-template/app/modules/user/service"
	"github.com/adamnasrudin03/go-template/pkg/driver"
	"github.com/sirupsen/logrus"
)

// Services all service object injected here
type Services struct {
	Auth authSrv.AuthService
	Msg  messageSrv.MessageService
	User userSrv.UserService
	Log  logSrv.LogService
}

func WiringService(repo *Repositories, cache *driver.RedisClient, cfg *configs.Configs, logger *logrus.Logger) *Services {

	return &Services{
		Auth: regAuthSrv(repo, cfg, logger),
		Msg:  regMsgSrv(repo, cfg, logger),
		User: regUserSrv(repo, cfg, logger),
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

func regMsgSrv(repo *Repositories, cfg *configs.Configs, logger *logrus.Logger) messageSrv.MessageService {
	return messageSrv.MessageService(messageSrv.MessageSrv{
		Repo:   repo.Message,
		Cfg:    cfg,
		Logger: logger,
	})
}

func regUserSrv(repo *Repositories, cfg *configs.Configs, logger *logrus.Logger) userSrv.UserService {
	params := userSrv.UserSrv{
		Repo:        repo.User,
		RepoCache:   repo.Cache,
		RepoLog:     repo.Log,
		RepoMessage: repo.Message,
		Cfg:         cfg,
		Logger:      logger,
	}
	return userSrv.NewUserService(params)
}
