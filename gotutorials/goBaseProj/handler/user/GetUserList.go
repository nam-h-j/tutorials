package user

import (
	"database/sql"

	"../../model"
	"./service"

	"github.com/gin-gonic/gin"
)

func GetUserList(c *gin.Context){

	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		panic(ok)
	}
	userService := service.UserService{db}

	resp := model.UserListResult{}
	resp = userService.GetUserList()
	c.JSON(resp.Status, resp)

}