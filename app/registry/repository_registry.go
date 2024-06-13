package registry

import (
	authRepo "github.com/adamnasrudin03/go-template/app/modules/auth/repository"
	logRepo "github.com/adamnasrudin03/go-template/app/modules/log/repository"
	userRepo "github.com/adamnasrudin03/go-template/app/modules/user/repository"
)

// Repositories all repo object injected here
type Repositories struct {
	Auth authRepo.AuthRepository
	User userRepo.UserRepository
	Log  logRepo.LogRepository
}
