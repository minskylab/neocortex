package neocortex

import (
	"net/http"
	"strconv"

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
		c.JSON(http.StatusOK, gin.H{
			"data": dialog,
		})
	})

	r.GET("/dialogs/*view", func(c *gin.Context) {
		frame := TimeFrame{}
		preset := TimeFramePreset(c.Query("preset"))

		if preset != DayPreset && preset != WeekPreset && preset != MonthPreset && preset != YearPreset {
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

			frame.From = from
			frame.To = to
		} else {
			frame.Preset = preset
		}

		viewID := c.Param("view")

		page, _ := strconv.Atoi(c.Query("page"))
		size, _ := strconv.Atoi(c.Query("size"))

		frame.PageNum = page
		frame.PageSize = size

		if viewID == "" || viewID == "/" {
			dialogs, err := api.repository.AllDialogs(frame)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"data": dialogs,
			})
			return
		}

		if _, err := xid.FromString(viewID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid view id"})
			return
		}

		dialogs, err := api.repository.DialogsByView(viewID, frame)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": dialogs,
		})
		return
	})
}
