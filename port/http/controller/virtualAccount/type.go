package controller_virtualAccount

import (
	"github.com/fbriansyah/pg-tools/port/http/controller"
)

type VirtualAccountController struct{}

func New() *VirtualAccountController {
	return &VirtualAccountController{}
}

var _ controller.IVirtualAccountController = (*VirtualAccountController)(nil)
