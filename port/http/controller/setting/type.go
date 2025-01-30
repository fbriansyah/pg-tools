package controller_setting

import "github.com/fbriansyah/pg-tools/port/http/controller"

type SettingController struct{}

func New() *SettingController {
	return &SettingController{}
}

var _ controller.ISettingController = (*SettingController)(nil)
