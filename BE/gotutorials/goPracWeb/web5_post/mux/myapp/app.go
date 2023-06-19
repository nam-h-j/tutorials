package myapp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

var userMap map[int]*User
var lastID int

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "this is index")
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "user, get user info : /user/{id}")
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	user, ok := userMap[id]
	if !ok {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No User ID : ", id)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(user)
	fmt.Fprint(w, string(data))
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	user := new(User) // create User struct
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	// Created User
	lastID++
	user.ID = lastID
	user.CreatedAt = time.Now()
	userMap[user.ID] = user

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	data, _ := json.Marshal(user) // 마샬링, 논리적 구조를 로우바이트로 변경하는 것(인코딩)
	fmt.Fprint(w, string(data))
}

// NewHandler
func NewHandler() http.Handler {
	userMap = make(map[int]*User) //init userMap
	lastID = 0
	mux := mux.NewRouter()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/user/{id:[0-9]+}", getUserHandler).Methods("GET")
	mux.HandleFunc("/user", userHandler).Methods("GET")
	mux.HandleFunc("/user", createUserHandler).Methods("POST")
	return mux
}
