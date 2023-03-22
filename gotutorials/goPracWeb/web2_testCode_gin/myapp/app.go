package myapp

import (
	"encoding/json"
	"fmt"
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
	// 3 way to get Query Params
	// name := c.Request.URL.Query().Get("name")
	//or
	// name := c.Query("name")
	//or with default values
	name := c.DefaultQuery("name", "Guest")

	c.String(http.StatusOK, "welcome %s!", name)
}

func postUserHandler(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Request: ", err)
		return
	}
	user.CreatedAt = time.Now()
	data, _ := json.Marshal(user)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, string(data))
}

func NewHttpHandler() *gin.Engine {
	router := gin.Default()

	router.GET("/", indexHandler)
	router.GET("/userName", getUserNameFromParam)
	router.POST("/user", postUserHandler)

	return router
}
