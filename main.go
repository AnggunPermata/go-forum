package main

import (
	"net/http"

	"github.com/AnggunPermata/go-forum/helper/utils"
	"github.com/AnggunPermata/go-forum/server/routes"
)

func init() {
	utils.LoadConfig()
}

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(utils.Config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	routes.InitRoutes(mux)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
