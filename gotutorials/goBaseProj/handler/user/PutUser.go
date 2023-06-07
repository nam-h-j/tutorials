package user

import (
	"database/sql"
	"net/http"

	"goBaseProj/handler/user/service"
	"goBaseProj/model"

	"github.com/gin-gonic/gin"
)

// PutUser godoc
// @Tags         유저 정보 관리
// @Summary      유저 정보 등록
// @Description  유저 정보를 등록합니다.
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

	res := model.UserResult{}
	res = userService.PutUser(resBody)

	// 정상적인 INSERT 아니면 상태를 RETURN
	if res.Status != http.StatusOK {
		c.JSON(res.Status, res)
		return
	}
}
