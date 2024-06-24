package registry

import (
	"github.com/adamnasrudin03/go-template/app/configs"
	authDelivery "github.com/adamnasrudin03/go-template/app/modules/auth/delivery"
	logDelivery "github.com/adamnasrudin03/go-template/app/modules/log/delivery"
	messageDelivery "github.com/adamnasrudin03/go-template/app/modules/message/delivery"
	userDelivery "github.com/adamnasrudin03/go-template/app/modules/user/delivery"
	"github.com/sirupsen/logrus"
)

// Deliveries all Controller object injected here
type Deliveries struct {
	Auth    authDelivery.AuthDelivery
	User    userDelivery.UserDelivery
	Log     logDelivery.LogDelivery
	Message messageDelivery.MessageDelivery
}

func WiringDelivery(srv *Services, cfg *configs.Configs, logger *logrus.Logger) *Deliveries {
	return &Deliveries{
		Auth:    authDelivery.NewAuthDelivery(srv.Auth, logger),
		User:    userDelivery.NewUserDelivery(srv.User, logger),
		Log:     logDelivery.NewLogDelivery(srv.Log, logger),
		Message: messageDelivery.NewMessageDelivery(srv.Msg, srv.Log, cfg, logger),
	}
}
