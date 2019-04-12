package neocortex

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

type API struct {
	Port       string
	repository Repository
	engine     *gin.Engine
	prefix     string
}

func (api *API) Launch() error {
	api.engine.GET(api.prefix+"/dialog/:id}", func(c *gin.Context) {
		id := c.Param("id")
		dialog, err := api.repository.GetDialogByID(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, dialog)
	})

	api.engine.GET(api.prefix+"/dialogs", func(c *gin.Context) {
		limit, _ := strconv.Atoi(c.Query("limit"))
		personID := c.Query("person")
		from, _ := time.Parse(time.RFC3339, c.Query("from"))
		until, _ := time.Parse(time.RFC3339, c.Query("until"))
		sessionID := c.Query("session")
		timezone := c.Query("timezone")
		batch, _ := strconv.Atoi(c.DefaultQuery("items", "10"))
		page, _ := strconv.Atoi(c.DefaultQuery("items", "1"))
		if page <= 0 {
			page = 1
		}
		if batch <= 0 {
			batch = 10
		}

		dialogs, err := api.repository.GetDialogs(DialogFilter{
			Limit:     int64(limit),
			PersonID:  personID,
			SessionID: sessionID,
			From:      from,
			Until:     until,
			Timezone:  timezone,
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		page -= 1
		length := len(dialogs)
		f := length * page / batch
		t := f + batch
		c.JSON(http.StatusOK, dialogs[f:t])
	})
	log.Println("API listening at " + api.Port)
	return api.engine.Run(api.Port)
}
