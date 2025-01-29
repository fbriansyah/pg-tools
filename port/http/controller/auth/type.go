package controller_auth

import (
	"github.com/fbriansyah/pg-tools/port/http/controller"
)

type AuthController struct{}

func New() *AuthController {
	return &AuthController{}
}

var _ controller.IAuthController = (*AuthController)(nil)
