package neocortex

import (
	"log"
	"net/http"

	"time"

	"github.com/araddon/dateparse"
	"github.com/gin-gonic/gin"
)

func (api *API) registerSummaryAPI(r *gin.RouterGroup) {
	r.GET("/summary", func(c *gin.Context) {
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

		fromScratch := time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)

		pastSummary, err := api.repository.Summary(TimeFrame{From: fromScratch, To: frame.From})
		summary, err := api.repository.Summary(frame)

		summary.RecurrentUsers = pastSummary.RecurrentUsers - summary.RecurrentUsers

		for timezone, rec := range pastSummary.UsersByTimezone {
			summary.UsersByTimezone[timezone] = UsersSummary{
				News:       rec.News - summary.UsersByTimezone[timezone].News,
				Recurrents: rec.Recurrents - summary.UsersByTimezone[timezone].Recurrents,
			}
		}

		log.Println(summary)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": summary,
		})
	})
}
