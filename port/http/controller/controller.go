package controller

import "net/http"

type IDashboardController interface {
	Index(w http.ResponseWriter, r *http.Request) error
}
