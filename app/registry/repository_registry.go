package registry

import (
	userRepo "github.com/adamnasrudin03/go-template/app/modules/user/repository"
)

// Repositories all repo object injected here
type Repositories struct {
	User userRepo.UserRepository
}
