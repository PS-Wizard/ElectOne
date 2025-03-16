package routes

import (
	"net/http"

	"github.com/PS-Wizard/Plugs/views/pages"
	"github.com/a-h/templ"
)

func HandleRoutes(mux *http.ServeMux) {
	mux.Handle("/", templ.Handler(pages.Home()))
	mux.Handle("/login", templ.Handler(pages.Authentication(true)))
	mux.Handle("/signup", templ.Handler(pages.Authentication(false)))
}
