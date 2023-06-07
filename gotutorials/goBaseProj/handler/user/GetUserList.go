package user

import (
	"database/sql"

	"goBaseProj/handler/user/service"
	"goBaseProj/model"

	"github.com/gin-gonic/gin"
)

// GetUserList godoc
// @Tags         유저 정보 관리
// @Summary      유저 정보 목록 열람
// @Description  유저 정보 목록을 가져옵니다.
// @Produce      json
// @Success      200  {array}  model.UserListResult
// @Router       /user/ [get]
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
