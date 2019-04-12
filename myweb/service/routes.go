package service

import (
	"log"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"GetAccount",
		"GET",
		"/account/{accountId}",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			log.Println("I Love You")
			w.Write([]byte("{\"result\": \"ok\"}"))
		},
	},
}
