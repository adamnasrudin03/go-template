package registry

import (
	userDelivery "github.com/adamnasrudin03/go-template/app/modules/user/delivery"
)

// Deliveries all Controller object injected here
type Deliveries struct {
	User userDelivery.UserDelivery
}
