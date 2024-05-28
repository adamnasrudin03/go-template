package registry

import (
	logSrv "github.com/adamnasrudin03/go-template/app/modules/log/service"
	userSrv "github.com/adamnasrudin03/go-template/app/modules/user/service"
)

// Services all service object injected here
type Services struct {
	User userSrv.UserService
	Log  logSrv.LogService
}
