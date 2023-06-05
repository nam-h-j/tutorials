package user

import (
	"database/sql"

	"goBaseProj/handler/user/service"
	"goBaseProj/model"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	userId := c.Param("id")

	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		panic(ok)
	}
	userService := service.UserService{db}

	resp := model.UserResult{}
	resp = userService.GetUser(userId)

	c.Header("Content-Type", "application/json")
	c.JSON(resp.Status, resp)
}
