package user

import (
	"database/sql"

	"goBaseProj/model"

	"goBaseProj/handler/user/service"

	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {
	userId := c.Param("id")

	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		panic(ok)
	}
	userService := service.UserService{db}

	resp := model.UserResult{}
	resp = userService.DeleteUser(userId)

	c.JSON(resp.Status, resp)
}
