package neocortex

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api *API) registerSummaryAPI(r *gin.RouterGroup) {
	r.GET("/summary", func(c *gin.Context) {
		summary, err := api.repository.Summary()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": summary,
		})
	})
}
