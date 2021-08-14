// Package webserver is used for REST endpoint requests. I'm really lazy to make a graphql endpoint
// If you want to add that, feel free to create a PR.
package webserver

import (
	"github.com/gorilla/mux"
	"github.com/provsalt/beast-backend/config"
	"github.com/provsalt/beast-backend/webserver/root"
	"github.com/provsalt/beast-backend/webserver/stats"
	"log"
	"net/http"
	"time"
)

func New(cfg config.Config) {
	r := mux.NewRouter()
	srv := &http.Server{
		Addr:         cfg.Webserver.Host,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	r.HandleFunc("/", root.Handler).Methods(http.MethodGet)
	r.HandleFunc("/stats", stats.QueryStats).Methods(http.MethodGet)

	r.Use(mux.CORSMethodMiddleware(r))

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
	select {}
}
