package controller_virtualAccount

import (
	"net/http"

	view_pages "github.com/fbriansyah/pg-tools/internal/view/pages"
	"github.com/fbriansyah/pg-tools/port/http/controller"
)

// Index implements controller.IVirtualAccountController.
func (v *VirtualAccountController) Index(w http.ResponseWriter, r *http.Request) error {
	return controller.Render(r, w, view_pages.VirtualAccountPage())
}
