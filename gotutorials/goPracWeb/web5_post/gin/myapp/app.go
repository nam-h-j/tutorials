package myapp

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
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

func indexHandler(c *gin.Context) {
	c.String(http.StatusOK, "this is index")
}

func userHandler(c *gin.Context) {
	c.String(http.StatusOK, "user, get user info : /user/{id}")
}

func getUserHandler(c *gin.Context) {
	idStr, _ := c.Params.Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	user, ok := userMap[id]
	if !ok {
		c.String(http.StatusOK, "No User ID : %v", id)
	}
	c.Writer.Header().Set("Content-Type", "application/json")
	data, _ := json.Marshal(user)
	c.String(http.StatusOK, string(data))
}

func createUserHandler(c *gin.Context) {
	user := new(User) // create User struct
	err := json.NewDecoder(c.Request.Body).Decode(user)
	if err != nil {
		c.String(http.StatusOK, err.Error())
		return
	}

	// Created User
	lastID++
	user.ID = lastID
	user.CreatedAt = time.Now()
	userMap[user.ID] = user

	c.Writer.Header().Set("Content-Type", "application/json")
	data, _ := json.Marshal(user) // 마샬링, 논리적 구조를 로우바이트로 변경하는 것(인코딩)
	c.String(http.StatusCreated, string(data))
}

// NewHandler
func NewHandler() *gin.Engine {
	userMap = make(map[int]*User) //init userMap
	lastID = 0
	router := gin.Default()

	router.GET("/", indexHandler)
	router.GET("/user", userHandler)
	router.GET("/user/:id", getUserHandler)
	router.POST("/user", createUserHandler)

	return router
}
