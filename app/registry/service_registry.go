package registry

import (
	authSrv "github.com/adamnasrudin03/go-template/app/modules/auth/service"
	logSrv "github.com/adamnasrudin03/go-template/app/modules/log/service"
	userSrv "github.com/adamnasrudin03/go-template/app/modules/user/service"
)

// Services all service object injected here
type Services struct {
	Auth authSrv.AuthService
	User userSrv.UserService
	Log  logSrv.LogService
}
