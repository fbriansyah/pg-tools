package controller_user

import "github.com/fbriansyah/pg-tools/port/http/controller"

type UserController struct{}

func New() *UserController {
	return &UserController{}
}

var _ controller.IUserController = (*UserController)(nil)
