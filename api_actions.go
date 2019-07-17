package neocortex

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api *API) registerActionsAPI(r *gin.RouterGroup) {

	r.POST("/actions/env/:name", func(c *gin.Context) {
		name := c.Param("name")

		type bind struct {
			Value string `json:"value"`
		}

		value := new(bind)

		err := c.BindJSON(value)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		err = api.repository.SetActionVar(name, value.Value)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": value,
		})
	})

	r.GET("/actions/env/:name", func(c *gin.Context) {
		name := c.Param("name")
		value, err := api.repository.GetActionVar(name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": value,
		})
	})
}
