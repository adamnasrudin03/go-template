package router

import (
	"net/http"

	"github.com/adamnasrudin03/go-template/app/registry"
	"github.com/adamnasrudin03/go-template/pkg/helpers"

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
		helpers.RenderJSON(c.Writer, http.StatusOK, "welcome this server")
	})

	v1 := r.router.Group("/api/v1")
	r.userRouter(v1, h.User)
	r.userRootRouter(v1, h.User)
	r.LogRouter(v1, h.Log)

	// // Static file source
	// staticFileRoutes := r.router.Group("/file")
	// staticFileRoutes.Use(middlewares.Authentication())
	// staticFileRoutes.StaticFS("/", http.Dir("public"))

	r.router.NoRoute(func(c *gin.Context) {
		err = helpers.ErrRouteNotFound()
		helpers.RenderJSON(c.Writer, http.StatusNotFound, err)
	})
	return r
}

func (r routes) Run(addr string) error {
	return r.router.Run(addr)
}
