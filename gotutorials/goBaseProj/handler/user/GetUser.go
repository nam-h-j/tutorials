package user

import (
	"database/sql"

	"../../model"
	"./service"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context){
	userId := c.Param("id")

	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		panic(ok)
	}
	userService := service.UserService{db}

	resp := model.UserResult{}
	resp = userService.GetUser(userId)
	c.JSON(resp.Status, resp)
}