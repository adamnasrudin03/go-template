package registry

import (
	logDelivery "github.com/adamnasrudin03/go-template/app/modules/log/delivery"
	userDelivery "github.com/adamnasrudin03/go-template/app/modules/user/delivery"
)

// Deliveries all Controller object injected here
type Deliveries struct {
	User userDelivery.UserDelivery
	Log  logDelivery.LogDelivery
}
