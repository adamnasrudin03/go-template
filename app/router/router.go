package router

import (
	"net/http"

	"github.com/adamnasrudin03/go-template/app/registry"

	response_mapper "github.com/adamnasrudin03/go-helpers/response-mapper/v1"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type routes struct {
	router *gin.Engine
}

func NewRoutes(h registry.Deliveries) routes {
	var err error
	r := routes{
		router: gin.Default(),
	}

	r.router.Use(gin.Logger())
	r.router.Use(gin.Recovery())
	r.router.Use(cors.Default())

	r.router.GET("/", func(c *gin.Context) {
		response_mapper.RenderJSON(c.Writer, http.StatusOK, "welcome this server")
	})

	v1 := r.router.Group("/api/v1")
	r.authRouter(v1, h.Auth)
	r.userRouter(v1, h.User)
	r.LogRouter(v1, h.Log)
	r.MessageRouter(v1, h.Message)

	// // Static file source
	// staticFileRoutes := r.router.Group("/file")
	// staticFileRoutes.Use(middlewares.Authentication())
	// staticFileRoutes.StaticFS("/", http.Dir("public"))

	r.router.NoRoute(func(c *gin.Context) {
		err = response_mapper.ErrRouteNotFound()
		response_mapper.RenderJSON(c.Writer, http.StatusNotFound, err)
	})
	return r
}

func (r routes) Run(addr string) error {
	return r.router.Run(addr)
}
