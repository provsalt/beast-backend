package webserver

import (
	"github.com/gorilla/mux"
	"github.com/provsalt/beast-backend/config"
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

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
}
