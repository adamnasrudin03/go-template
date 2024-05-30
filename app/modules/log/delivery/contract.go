package delivery

import (
	"github.com/adamnasrudin03/go-template/app/modules/log/service"
	"github.com/gin-gonic/gin"
)

type LogDelivery interface {
	GetList(ctx *gin.Context)
}

type logDel struct {
	Service service.LogService
}

func NewLogDelivery(srv service.LogService) LogDelivery {
	return &logDel{
		Service: srv,
	}
}
