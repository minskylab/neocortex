package neocortex

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type API struct {
	Port       string
	repository Repository
	prefix     string
}

func (api *API) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		log.Println(req.URL.Path)
		if strings.HasPrefix(req.URL.Path, api.prefix+"/dialog/") {
			var id string
			_, err := fmt.Sscanf(req.URL.Path, api.prefix+"/dialog/%s", &id)
			if err != nil {
				http.Error(res, err.Error(), http.StatusInternalServerError)
				return
			}
			id = strings.Replace(id, "/", "", -1)
			dialog, err := api.repository.GetDialogByID(id)
			if err != nil {
				http.Error(res, err.Error(), http.StatusInternalServerError)
				return
			}
			data, err := json.Marshal(dialog)
			if err != nil {
				http.Error(res, err.Error(), http.StatusInternalServerError)
				return
			}
			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(http.StatusOK)
			_, err = io.WriteString(res, string(data))
			if err != nil {
				http.Error(res, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		} else if strings.HasPrefix(req.URL.Path, api.prefix+"/dialogs") {
			limit, _ := strconv.Atoi(req.URL.Query().Get("limit"))
			personID := req.URL.Query().Get("person")
			from, _ := time.Parse(time.RFC3339, req.URL.Query().Get("from"))
			until, _ := time.Parse(time.RFC3339, req.URL.Query().Get("until"))
			sessionID := req.URL.Query().Get("session")
			timezone := req.URL.Query().Get("timezone")
			batch, _ := strconv.Atoi(req.URL.Query().Get("items"))
			page, _ := strconv.Atoi(req.URL.Query().Get("page"))
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
				http.Error(res, err.Error(), http.StatusInternalServerError)
				return
			}

			page -= 1
			length := len(dialogs)
			f := length * page / batch
			t := f + batch
			if t > length {
				t = length
			}
			data, err := json.Marshal(dialogs[f:t])
			if err != nil {
				http.Error(res, err.Error(), http.StatusInternalServerError)
				return
			}
			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(http.StatusOK)
			_, err = io.WriteString(res, string(data))
			if err != nil {
				http.Error(res, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}
	}

	http.Error(res, "Not Found", http.StatusNotFound)
}

func (api *API) Launch() error {
	log.Println("API listening at " + api.Port)
	return http.ListenAndServe(api.Port, api)
}
