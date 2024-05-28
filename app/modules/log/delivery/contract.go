package delivery

import (
	"github.com/adamnasrudin03/go-template/app/modules/log/service"
)

type LogDelivery interface {
}

type logDel struct {
	Service service.LogService
}

func NewLogDelivery(srv service.LogService) LogDelivery {
	return &logDel{
		Service: srv,
	}
}
