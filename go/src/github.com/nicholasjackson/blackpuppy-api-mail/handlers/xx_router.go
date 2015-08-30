package handlers

import (
	"github.com/gorilla/pat"
)

func GetRouter() *pat.Router {
	r := pat.New()

	r.Get("/mail/healthcheck", HelloWorldHandler)
	r.Post("/mail/contactus", ContactUsHandler)
	return r
}
