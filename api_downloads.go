package neocortex

import (
	"bytes"
	"encoding/csv"
	"net/http"

	"strconv"

	"io/ioutil"
	"os"

	"github.com/araddon/dateparse"
	"github.com/gin-gonic/gin"
)

func (api *API) registerDownloadsAPI(r *gin.RouterGroup) {
	r.POST("/download/chat/:userID", func(c *gin.Context) {
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

		userID := c.Param("userID")

		if userID == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user ID, please select one or all"})
			return
		}

		chatsFound := make([]*Chat, 0)

		if userID == "all" {
			timezone := c.Query("timezone")
			if timezone == "" {
				chatsFound = chats
			} else {
				for _, c := range chats {
					if c.Person.Timezone == timezone {
						chatsFound = append(chatsFound, c)
					}
				}
			}
		} else {
			for _, c := range chats {
				if c.Person.ID == userID {
					chatsFound = append(chatsFound, c)
				}
			}
		}

		if len(chatsFound) == 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "user ID not found in your time frame"})
			return
		}

		table := [][]string{{
			"id",
			"name",
			"timezone",
			"last_message",
			"performance",
			"message_at",
			"message_who",
			"message_message",
			"message_intent",
			"message_entities",
		}}

		lastID := ""
		row := []string{}
		for _, c := range chatsFound {
			if c.ID != lastID {
				if len(c.Messages) > 0 {
					r, ok := c.Messages[0].Response.Value.(string)
					if !ok {
						r = ""
					}
					row = []string{
						c.ID,
						c.Person.Name,
						c.Person.Timezone,
						c.LastMessageAt.String(),
						strconv.FormatFloat(c.Performance, 'f', 10, 64),
						c.Messages[0].At.String(),
						string(c.Messages[0].Owner),
						r,
						"",
						"",
					}
				}
				lastID = c.ID
				table = append(table, row)
				row = []string{"", "", "", "", ""}
			}

			if len(c.Messages) > 1 {
				for _, m := range c.Messages[1:] {
					row = []string{"", "", "", "", ""}
					row = append(row, m.At.String())
					row = append(row, string(m.Owner))

					if m.Response.Type == "text" {
						r, ok := m.Response.Value.(string)
						if !ok {
							r = ""
						}
						row = append(row, r)
					} else {
						row = append(row, "")
					}

					intent := ""

					if m.Intents != nil {
						if len(m.Intents) > 0 {
							intent = m.Intents[0].Intent
						}
					}

					row = append(row, intent)

					entity := ""
					if m.Entities != nil {
						ents := ""
						for _, e := range m.Entities {
							ents = ents + "|" + e.Entity
						}
						entity = ents
					}
					row = append(row, entity)

					table = append(table, row)
				}
			}
		}

		file := bytes.NewBufferString("")
		err = csv.NewWriter(file).WriteAll(table)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		tempFile, err := ioutil.TempFile(os.TempDir(), "neo")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		_, err = tempFile.Write(file.Bytes())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Status(http.StatusOK)
		c.Header("Content-Type", "text/csv")
		c.File(tempFile.Name())
	})
}
