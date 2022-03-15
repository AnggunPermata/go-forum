package routes

import (
	"net/http"

	"github.com/AnggunPermata/go-forum/server/handler"
)

func InitRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", handler.Index)
}
