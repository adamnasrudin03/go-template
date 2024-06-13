package registry

import (
	authDelivery "github.com/adamnasrudin03/go-template/app/modules/auth/delivery"
	logDelivery "github.com/adamnasrudin03/go-template/app/modules/log/delivery"
	messageDelivery "github.com/adamnasrudin03/go-template/app/modules/message/delivery"
	userDelivery "github.com/adamnasrudin03/go-template/app/modules/user/delivery"
)

// Deliveries all Controller object injected here
type Deliveries struct {
	Auth    authDelivery.AuthDelivery
	User    userDelivery.UserDelivery
	Log     logDelivery.LogDelivery
	Message messageDelivery.MessageDelivery
}
