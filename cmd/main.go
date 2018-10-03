package main

import (
	"api-boilerplate/router"
	"net/http"
)

func main() {
	r := router.NewRouter()
	http.ListenAndServe(":3030", r)
}
