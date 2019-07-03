package neocortex

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api *API) registerViewsAPI(r *gin.RouterGroup) {
	r.POST("/view", func(c *gin.Context) {
		view := new(View)

		if err := c.BindJSON(view); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err := api.repository.SaveView(view); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": view})
	})

	r.GET("/view/:id", func(c *gin.Context) {
		id := c.Param("id")
		view, err := api.repository.GetViewByID(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": view})
	})

	r.GET("/views/*name", func(c *gin.Context) {
		name := c.Param("name")
		if name == "" || name == "/" {
			views, err := api.repository.AllViews()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"data": views})
			return
		}

		views, err := api.repository.FindViewByName(name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": views})
	})

	r.PUT("/view/:id", func(c *gin.Context) {
		view := new(View)

		if err := c.BindJSON(view); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		err := api.repository.UpdateView(view)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": view})
	})
}
