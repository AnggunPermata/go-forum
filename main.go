package main

import (
	"net/http"
	"time"

	"github.com/AnggunPermata/go-forum/helper/utils"
	"github.com/AnggunPermata/go-forum/server/routes"
)

func init() {
	utils.LoadConfig()
	utils.InitLogger()
}

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(utils.Config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	routes.InitRoutes(mux)

	server := &http.Server{
		Addr:           utils.Config.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(utils.Config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(utils.Config.WriteTimeOut * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}
