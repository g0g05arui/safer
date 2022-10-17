package main

import (
	"log"
	"net/http"

	"safer.com/m/api/v1/router"
	"safer.com/m/internal/env"
	_ "safer.com/m/services"
)

func main() {

	r := router.Init()
	err := http.ListenAndServe(env.Cfg["PORT"], r)
	if err != nil {
		log.Fatal(err)
	}
}
