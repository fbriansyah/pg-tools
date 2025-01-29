package controller_auth

import (
	"net/http"

	view_auth "github.com/fbriansyah/pg-tools/internal/view/auth"
	"github.com/fbriansyah/pg-tools/port/http/controller"
)

// SignIn implements controller.IAuthController.
func (a *AuthController) SignIn(w http.ResponseWriter, r *http.Request) error {
	return controller.Render(r, w, view_auth.SignIn())
}
