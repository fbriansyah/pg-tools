package controller_dashboard

import (
	"github.com/fbriansyah/pg-tools/port/http/controller"
)

type DashboardController struct{}

func New() *DashboardController {
	return &DashboardController{}
}

var _ controller.IDashboardController = &DashboardController{}
