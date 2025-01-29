package http

import (
	"embed"
	"net/http"
	"time"

	"github.com/fbriansyah/pg-tools/port/http/controller"
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

type RouterModule struct {
	DashboardCtrl controller.IDashboardController
}

func NewRoute(FS *embed.FS, rm *RouterModule) http.Handler {
	requestTimeout := time.Duration(60) * time.Second
	router := chi.NewRouter()

	// Init default middleware from chi
	router.Use(chiMiddleware.RealIP)
	router.Use(chiMiddleware.Recoverer)
	router.Use(chiMiddleware.Timeout(requestTimeout))

	router.Handle("/*", http.StripPrefix("/", http.FileServer(http.FS(FS))))
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/dashboard", http.StatusFound)
	})
	router.Get("/dashboard", controller.MakeHandler(rm.DashboardCtrl.Index))

	return router
}
