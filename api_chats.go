package neocortex

import (
	"net/http"
	"strconv"

	"github.com/araddon/dateparse"
	"github.com/gin-gonic/gin"
)

func (api *API) registerChatsAPI(r *gin.RouterGroup) {
	r.GET("/chat/:id", func(c *gin.Context) {
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

		page, _ := strconv.Atoi(c.Query("page"))
		size, _ := strconv.Atoi(c.Query("size"))

		frame.PageNum = page
		frame.PageSize = size

		dialogs, err := api.repository.AllDialogs(frame)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		chats := api.analytics.processDialogs(dialogs)

		userID := c.Param("id")

		if userID == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user ID, please select one or all"})
			return
		}

		if userID == "all" {
			c.JSON(http.StatusOK, gin.H{
				"data": chats,
			})
			return
		}

		chatFound := new(Chat)
		for _, c := range chats {
			if c.Person.ID == userID {
				chatFound = c
			}
		}

		if chatFound == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "user ID not found in your time frame"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": chatFound,
		})
	})

	r.GET("/chats", func(c *gin.Context) {
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

		page, _ := strconv.Atoi(c.Query("page"))
		size, _ := strconv.Atoi(c.Query("size"))

		frame.PageNum = page
		frame.PageSize = size

		dialogs, err := api.repository.AllDialogs(frame)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		chats := api.analytics.processDialogs(dialogs)
		c.JSON(http.StatusOK, gin.H{
			"data": chats,
		})
	})

}
