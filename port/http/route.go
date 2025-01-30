package http

import (
	"embed"
	"log/slog"
	"net/http"
	"time"

	"github.com/fbriansyah/pg-tools/port/http/controller"
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

type RouterModule struct {
	DashboardCtrl controller.IDashboardController
	AuthCtrl      controller.IAuthController
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
	router.Route("/auth", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				http.Redirect(w, r, "/auth/sign-in", http.StatusFound)
			})
			r.Get("/sign-in", controller.MakeHandler(rm.AuthCtrl.SignIn))
			r.Post("/sign-in", func(w http.ResponseWriter, r *http.Request) {
				email := r.FormValue("email")
				password := r.FormValue("password")
				slog.Info("Sing in Info", "email", email, "pass", password)
				http.Redirect(w, r, "/dashboard", http.StatusFound)
			})
		})
	})

	return router
}
