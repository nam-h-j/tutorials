package user

import (
	"database/sql"
	"net/http"

	"../../model"
	"./service"

	"github.com/gin-gonic/gin"
)

func PostUser(c *gin.Context) {
	reqBodyUser := model.User{}
	
	// 미들웨어의 디비풀로 연결
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		panic(ok)
	}

	// 유저 객체 
	if error := c.BindJSON(&reqBodyUser); error != nil {
		return
	}

	// err := c.BindJson(&paramUser)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	// userService 호출
	userService := service.UserService{db}
	// 쿼리날리기
	insertRes := userService.PostUser(reqBodyUser)

	// 정상적인 INSERT 아니면 상태를 RETURN
	if insertRes.Status != http.StatusOK {
		c.JSON(insertRes.Status, insertRes)
		return
	}

	// 성공했다면 insert된 값을 반환
	// resp := userService.GetByAccountSrl(insertRes.Message)
	// resp.Cmd = "INSERT"
	// c.JSON(resp.Status, resp)

}
