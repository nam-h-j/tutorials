package main

import (
	"net/http" 
	"./myapp"
	)

func NewHandler() http.Handler(){
	mux := myapp.NewHandler()
	return mux
}

func main() {
	mux := NewHandler()
	http.ListenAndServe(":1234", mux)
}
