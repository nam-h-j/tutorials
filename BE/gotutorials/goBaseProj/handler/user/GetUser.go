package user

import (
	"database/sql"
	"fmt"
	"strconv"

	"goBaseProj/handler/user/service"
	"goBaseProj/model"

	"github.com/gin-gonic/gin"
)

// GetUser godoc
// @Tags         유저 정보 관리
// @Summary      유저 정보 열람
// @Description  user_id과 관련된 유저 정보값을 가져옵니다.
// @Produce      json
// @Param 		   userId	path int true "가져올 유저 정보의 유저 시리얼(아이디)"
// @Success      200  {array}  model.UserResult
// @Router       /user/{userId} [get]
func GetUser(c *gin.Context) {
	userIdParam := c.Param("id")
	userId, err := strconv.Atoi(userIdParam)
	if err != nil {
		fmt.Println(err)
	}
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
