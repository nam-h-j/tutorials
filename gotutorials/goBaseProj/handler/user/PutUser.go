package user

import (
	"database/sql"

	"goBaseProj/handler/user/service"
	"goBaseProj/model"

	"github.com/gin-gonic/gin"
)

// PutUser godoc
// @Tags         유저 정보 관리
// @Summary      유저 정보 수정
// @Description  유저 정보를 수정합니다.
// @Produce      json
// @Param 		 Param body model.User true "유저 정보 JSON Format"
// @Success      200  {array}  model.UserResult
// @Router       /user [put]
func PutUser(c *gin.Context) {
	resBody := model.User{}

	// 디비풀 연결
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		panic(ok)
	}

	// response 객체 검사
	err := c.BindJSON(&resBody)
	if err != nil {
		return
	}

	// userService 호출
	userService := service.UserService{db}

	resp := model.UserResult{}
	resp = userService.PutUser(resBody)

	c.Header("Content-Type", "application/json")
	c.JSON(resp.Status, resp)
}
