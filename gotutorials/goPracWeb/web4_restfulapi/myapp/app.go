package myapp

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "this is index")
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "user, get user info : /user/{id}")
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, "user id: ", vars["id"])
}

// NewHandler
func NewHandler() http.Handler {
	mux := mux.NewRouter()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/user", userHandler)
	mux.HandleFunc("/user/{id:[0-9]+}", getUserHandler)
	return mux
}
