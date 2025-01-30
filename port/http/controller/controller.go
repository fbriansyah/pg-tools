package controller

import "net/http"

type IDashboardController interface {
	Index(w http.ResponseWriter, r *http.Request) error
}

type IAuthController interface {
	SignIn(w http.ResponseWriter, r *http.Request) error
}

type IVirtualAccountController interface {
	Index(w http.ResponseWriter, r *http.Request) error
}

type IQrisController interface {
	Index(w http.ResponseWriter, r *http.Request) error
}

type IUserController interface {
	Index(w http.ResponseWriter, r *http.Request) error
}

type ISettingController interface {
	Index(w http.ResponseWriter, r *http.Request) error
}
