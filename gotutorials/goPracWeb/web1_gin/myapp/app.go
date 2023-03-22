package myapp

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type User struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func userHandler(c *gin.Context) {
	user := new(User)
	err := c.BindJSON(&user)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad Request: ", err.Error())
		return
	}
	user.CreatedAt = time.Now()
	data, _ := json.Marshal(user)
	c.IndentedJSON(http.StatusCreated, data)
}

func NewHttpHandler() {
	serve := gin.Default()
	serve.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "myappRootOk")
	})
	serve.GET("/user", userHandler)
	serve.Run(":1234")
}
