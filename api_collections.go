package neocortex

import (
	"net/http"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

func (api *API) registerCollectionsAPI(r *gin.RouterGroup) {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("/login", getJWTAuth().LoginHandler)

	auth := r.Group("/auth")

	auth.GET("/refresh_token", getJWTAuth().RefreshHandler)
	auth.Use(getJWTAuth().MiddlewareFunc())

	auth.GET("/collections/:type", func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		t := c.Param("type")

		if t != "all" {
			switch t {
			case "entity", "entities":
				ents := api.repository.Entities()
				c.JSON(http.StatusOK, gin.H{
					"userID": claims["id"],
					"data":   ents,
				})
				return
			case "intent", "intents":
				ints := api.repository.Intents()
				c.JSON(http.StatusOK, gin.H{
					"userID": claims["id"],
					"data":   ints,
				})
				return
			case "node", "dialog", "nodes", "dialog_nodes", "dialogs":
				nodes := api.repository.DialogNodes()
				c.JSON(http.StatusOK, gin.H{
					"userID": claims["id"],
					"data":   nodes,
				})
				return
			case "context_vars", "vars":
				vars := api.repository.ContextVars()
				c.JSON(http.StatusOK, gin.H{
					"userID": claims["id"],
					"data":   vars,
				})
				return
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid type of collection"})
				return
			}
		}

		ents := api.repository.Entities()
		ints := api.repository.Intents()
		nodes := api.repository.DialogNodes()
		vars := api.repository.ContextVars()

		c.JSON(http.StatusOK, gin.H{"data": gin.H{
			"userID":       claims["id"],
			"intents":      ints,
			"entities":     ents,
			"nodes":        nodes,
			"context_vars": vars,
		}})
	})
}
