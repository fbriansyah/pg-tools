package controller_qris

import "github.com/fbriansyah/pg-tools/port/http/controller"

type QrisController struct{}

func New() *QrisController {
	return &QrisController{}
}

var _ controller.IQrisController = (*QrisController)(nil)
