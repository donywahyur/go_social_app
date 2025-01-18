package main

import (
	"go_social_app/internal/database"
	"go_social_app/internal/delivery/rest"
	"go_social_app/internal/env"
	"log"
	"net/http"
	"time"
)

type application struct {
	config config
}

type config struct {
	addr string
}

func main() {
	cfg := config{
		addr: env.Get("ADDR", ":8080"),
	}
	app := &application{
		config: cfg,
	}
	db := database.GetDB()
	mux := rest.LoadRoutes(db)
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 30,
		IdleTimeout:  time.Second * 60,
	}
	log.Printf("server has started at %s", app.config.addr)
	log.Fatal(srv.ListenAndServe())
}
