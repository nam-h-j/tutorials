package myapp

import (
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

func indexHandler(c *gin.Context) {
	c.String(http.StatusOK, "myapp_index")
}

func getUserNameFromParam(c *gin.Context) {
	name := c.DefaultQuery("name", "Guest")
	c.String(http.StatusOK, "welcome %s!", name)
}

func NewHttpHandler() *gin.Engine {
	router := gin.Default()

	router.GET("/", indexHandler)
	router.GET("/user", getUserNameFromParam)

	return router
}
