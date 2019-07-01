package neocortex

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api *API) registerCollectionsAPI(r *gin.RouterGroup) {
	r.GET("/collections/:type", func(c *gin.Context) {
		t := c.Param("type")

		if t != "" {
			switch t {
			case "entity", "entities":
				ents := api.repository.Entities()
				c.JSON(http.StatusOK, gin.H{"data": ents})
				return
			case "intent", "intents":
				ints := api.repository.Intents()
				c.JSON(http.StatusOK, gin.H{"data": ints})
				return
			case "node", "dialog", "nodes", "dialog_nodes", "dialogs":
				nodes := api.repository.DialogNodes()
				c.JSON(http.StatusOK, gin.H{"data": nodes})
				return
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid type of collection"})
				return
			}
		}

		ents := api.repository.Entities()
		ints := api.repository.Intents()
		nodes := api.repository.DialogNodes()

		c.JSON(http.StatusOK, gin.H{"data": gin.H{
			"intents":  ints,
			"entities": ents,
			"nodes":    nodes,
		}})
	})
}
