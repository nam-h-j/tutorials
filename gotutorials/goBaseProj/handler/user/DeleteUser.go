package user

import (
	"database/sql"

	"goBaseProj/model"

	"goBaseProj/handler/user/service"

	"github.com/gin-gonic/gin"
)

// DeleteUser godoc
// @Tags         유저 정보 관리
// @Summary      유저 정보 열람
// @Description  user_id과 관련된 유저 정보값을 가져옵니다.
// @Produce      json
// @Param 		 userId	path	int		true  "삭제할 유저 정보의 유저 시리얼(아이디)"
// @Success      200  {array}  model.UserResult
// @Router       /user/{userId} [Delete]
func DeleteUser(c *gin.Context) {
	userId := c.Param("id")

	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		panic(ok)
	}
	userService := service.UserService{db}

	resp := model.UserResult{}
	resp = userService.DeleteUser(userId)

	c.Header("Content-Type", "application/json")
	c.JSON(resp.Status, resp)
}
