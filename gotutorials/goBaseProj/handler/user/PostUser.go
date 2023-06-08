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
// @Param 		 Param body model.UserPostReq true "유저 정보 JSON Format"
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
	postRes := userService.PostUser(resBody)

	// 성공했다면 insert된 값을 반환
	if postRes.Status == http.StatusOK {
		resp := userService.GetUser(postRes.UserData.UserID)
		resp.Cmd = "INSERT"
		c.JSON(resp.Status, resp)
	}
}
