package neocortex

import (
	"net/http"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

func (api *API) registerViewsAPI(r *gin.RouterGroup) {

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("/login", getJWTAuth().LoginHandler)

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

		c.JSON(http.StatusOK, gin.H{
			"data": view,
		})
	})

	auth := r.Group("/auth")

	auth.GET("/refresh_token", getJWTAuth().RefreshHandler)
	auth.Use(getJWTAuth().MiddlewareFunc())

	auth.GET("/view/:id", func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		id := c.Param("id")
		view, err := api.repository.GetViewByID(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"userID": claims["id"],
			"data":   view,
		})
	})

	auth.GET("/views/*name", func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		name := c.Param("name")
		if name == "" || name == "/" {
			views, err := api.repository.AllViews()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"userID": claims["id"],
				"data":   views,
			})
			return
		}

		views, err := api.repository.FindViewByName(name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"userID": claims["id"],
			"data":   views,
		})
	})

}
