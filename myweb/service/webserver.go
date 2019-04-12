package service

import (
	"log"
	"net/http"
)

func StartWebServer(port string) {
	r := NewRouter()
	http.Handle("/", r)
	log.Println("Start Http Server at http://127.0.0.1:" + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Println("error " + port)
		log.Println(err.Error())
	}
}
