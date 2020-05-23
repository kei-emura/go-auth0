package home

import (
	"go-auth0/routes/templates"
	"net/http"
)

// Handler is to render home.html
func Handler(w http.ResponseWriter, r *http.Request) {
	templates.RenderTemplate(w, "home", nil)
}
