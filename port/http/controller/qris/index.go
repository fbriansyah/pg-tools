package controller_qris

import (
	"net/http"

	view_pages "github.com/fbriansyah/pg-tools/internal/view/pages"
	"github.com/fbriansyah/pg-tools/port/http/controller"
)

// Index implements controller.IDashboardController.
func (d *QrisController) Index(w http.ResponseWriter, r *http.Request) error {
	return controller.Render(r, w, view_pages.QrisPage())
}
