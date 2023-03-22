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
	// 3 way to get Query Params in GIN
	// 1. name := c.Request.URL.Query().Get("name")
	// or
	// 2. name := c.Query("name")
	// or with default values
	// 3. name := c.DefaultQuery("name", "Guest")
	name := c.DefaultQuery("name", "Guest")
	c.String(http.StatusOK, "welcome %s!", name)
}

func NewHttpHandler() *gin.Engine {
	router := gin.Default()

	router.GET("/", indexHandler)
	router.GET("/user", getUserNameFromParam)

	return router
}
