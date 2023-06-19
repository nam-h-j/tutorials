package main

import (
	"net/http"

	"./myapp"
)

func main() {
	http.ListenAndServe(":1234", myapp.NewHandler())
}
