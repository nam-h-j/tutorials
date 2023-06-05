package user

import (
	"database/sql"

	"goBaseProj/handler/user/service"
	"goBaseProj/model"

	"github.com/gin-gonic/gin"
)

func GetUserList(c *gin.Context) {

	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		panic(ok)
	}
	userService := service.UserService{db}

	resp := model.UserListResult{}
	resp = userService.GetUserList()
	c.JSON(resp.Status, resp)

}
