package myapp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context) {
	c.String(http.StatusOK, "this is index")
}

func userHandler(c *gin.Context) {
	c.String(http.StatusOK, "user, get user info : /user/{id}")
}

func getUserHandler(c *gin.Context) {
	id, _ := c.Params.Get("id")
	c.String(http.StatusOK, "user id: %s", id)
}

func NewHandler() *gin.Engine {
	router := gin.Default()

	router.GET("/", indexHandler)
	router.GET("/user", userHandler)
	router.GET("/user/:id", getUserHandler)

	return router
}
