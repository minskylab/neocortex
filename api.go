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

func (api *API) registerEndpoints() {
	r := api.e.Group(api.prefix)
	api.registerDialogsAPI(r)
	api.registerViewsAPI(r)

}

func (api *API) Launch() error {
	api.registerEndpoints()
	return api.e.Run(api.Port)
}
