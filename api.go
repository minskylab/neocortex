package neocortex

import (
	"github.com/gin-gonic/gin"
)

type API struct {
	e          *gin.Engine
	Port       string
	repository Repository
	prefix     string
}

func newCortexAPI(repo Repository, prefix, port string) *API {
	return &API{
		e:          gin.Default(),
		Port:       port,
		prefix:     prefix,
		repository: repo,
	}
}

func (api *API) registerEndpoints(engine *Engine) {
	api.e.Use(gin.Logger())
	api.e.Use(gin.Recovery())
	authJWTMiddleware := getJWTAuth(engine)
	api.e.POST("/login", authJWTMiddleware.LoginHandler)
	api.e.GET("/token_refresh", authJWTMiddleware.RefreshHandler)
	api.e.Use(authJWTMiddleware.MiddlewareFunc())

	r := api.e.Group(api.prefix)
	api.registerDialogsAPI(r)
	api.registerViewsAPI(r)
	api.registerActionsAPI(r)
	api.registerCollectionsAPI(r)
}

func (api *API) Launch(engine *Engine) error {
	api.registerEndpoints(engine)
	return api.e.Run(api.Port)
}
