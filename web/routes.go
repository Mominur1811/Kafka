package web

import (
	"kafka-by-momin/web/handlers"
	"kafka-by-momin/web/middlewares"
	"net/http"
)

func InitRoutes(mux *http.ServeMux, manager *middlewares.Manager) {

	mux.Handle(
		"POST /insert",
		manager.With(
			http.HandlerFunc(handlers.CreateNotification),
		),
	)
}
