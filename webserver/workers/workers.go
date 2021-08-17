package workers

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

// QueryWorkers is a handler which just queries directly from the API for workers
func QueryWorkers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	address := r.URL.Query().Get("address")

	// basic address validation I guess? I'm too lazy to write one again
	if len(address) == 95 || len(address) == 105 {
		cookie, _ := cookiejar.New(nil)
		url, _ := url.Parse("https://xmrvsbeast.com")
		cookie.SetCookies(url, []*http.Cookie{&http.Cookie{Name: "wa", Value: address}})
		c := http.Client{Timeout: 5 * time.Second, Jar: cookie}
		response, err := c.Get("https://xmrvsbeast.com/workers/")
		if err != nil {
			log.Print(err.Error())
			return
		}

		var stats interface{}
		err = json.NewDecoder(response.Body).Decode(&stats)
		if err != nil {
			log.Print(err.Error())
			return
		}
		err = json.NewEncoder(w).Encode(&stats)
		if err != nil {
			log.Print(err.Error())
			return
		}
	} else {
		c := http.Client{Timeout: 5 * time.Second}
		response, err := c.Get("https://xmrvsbeast.com/workers/")
		if err != nil {
			log.Print(err.Error())
			return
		}

		var stats interface{}
		err = json.NewDecoder(response.Body).Decode(&stats)
		if err != nil {
			log.Print(err.Error())
			return
		}
		err = json.NewEncoder(w).Encode(&stats)
		if err != nil {
			log.Print(err.Error())
			return
		}
	}
}
