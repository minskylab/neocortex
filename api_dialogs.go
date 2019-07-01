package neocortex

import (
	"net/http"

	"github.com/araddon/dateparse"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

func (api *API) registerDialogsAPI(r *gin.RouterGroup) {
	r.GET("/dialog/:id", func(c *gin.Context) {
		id := c.Param("id")
		dialog, err := api.repository.GetDialogByID(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": dialog})
	})

	r.GET("/dialogs/*view", func(c *gin.Context) {
		from, err := dateparse.ParseAny(c.Query("from"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "from date not found"})
			return
		}

		to, err := dateparse.ParseAny(c.Query("to"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "to date not found"})
			return
		}

		viewID := c.Param("view")

		if viewID == "" {
			dialogs, err := api.repository.AllDialogs(TimeFrame{
				From: from,
				To:   to,
			})
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{"data": dialogs})
			return
		}

		if _, err := xid.FromString(viewID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid view id"})
			return
		}

		dialogs, err := api.repository.DialogsByView(viewID, TimeFrame{
			From: from,
			To:   to,
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": dialogs})
		return
	})

}
