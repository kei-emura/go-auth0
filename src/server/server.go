package server

import (
	"go-auth0/routes/callback"
	"go-auth0/routes/home"
	"go-auth0/routes/login"
	"go-auth0/routes/logout"
	"go-auth0/routes/middlewares"
	"go-auth0/routes/user"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// Run is to running server
func Run() error {
	r := mux.NewRouter()

	r.HandleFunc("/", home.Handler)
	r.HandleFunc("/login", login.Handler)
	r.HandleFunc("/logout", logout.Handler)
	r.HandleFunc("/callback", callback.Handler)
	r.Handle("/user", negroni.New(
		negroni.HandlerFunc(middlewares.IsAuthenticated),
		negroni.Wrap(http.HandlerFunc(user.Handler)),
	))
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))
	http.Handle("/", r)
	log.Print("Server listening on http://locahost:3000/")
	return http.ListenAndServe("0.0.0.0:3000", nil)
}
