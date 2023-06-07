package user

import (
	"database/sql"
	"net/http"

	"goBaseProj/handler/user/service"
	"goBaseProj/model"

	"github.com/gin-gonic/gin"
)

// PostUser godoc
// @Tags         유저 정보 관리
// @Summary      유저 정보 등록
// @Description  유저 정보를 등록합니다.
// @Produce      json
// @Param 		 Param body model.User true "유저 정보 JSON Format"
// @Success      200  {array}  model.UserResult
// @Router       /user [post]
func PostUser(c *gin.Context) {
	resBody := model.User{}

	// 미들웨어의 디비풀로 연결
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		panic(ok)
	}

	// 유저 객체
	if error := c.BindJSON(&resBody); error != nil {
		return
	}

	// userService 호출
	userService := service.UserService{db}
	// 쿼리날리기
	insertRes := userService.PostUser(resBody)

	// 정상적인 INSERT 아니면 상태를 RETURN
	if insertRes.Status != http.StatusOK {
		c.JSON(insertRes.Status, insertRes)
		return
	}

	// 성공했다면 insert된 값을 반환
	// resp := userService.GetUser(insertRes.Message)
	// resp.Cmd = "INSERT"
	// c.JSON(resp.Status, resp)
}
