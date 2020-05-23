package user

import (
	"go-auth0/app"
	"go-auth0/routes/templates"
	"net/http"
)

// Handler is to access user profile
func Handler(w http.ResponseWriter, r *http.Request) {
	session, err := app.Store.Get(r, "auth-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	templates.RenderTemplate(w, "user", session.Values["profile"])
}
