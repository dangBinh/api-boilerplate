package main

import (
	"api-boilerplate/router"
	"fmt"
	"net/http"
)

func main() {
	r := router.NewRouter()
	fmt.Println("Listening on port 3030")
	http.ListenAndServe(":3030", r)
}
