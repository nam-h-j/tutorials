package user

import (
	"database/sql"
	"net/http"

	"../../model"
	"./service"

	"github.com/gin-gonic/gin"
)

func PostUser(c *gin.Context) {
	paramUser := model.User{}

	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		panic(ok)
	}

	if error := c.BindJSON(&paramUser); error != nil {
		return
	}

	// err := c.BindJson(&paramUser)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	userService := service.UserService{db}
	insertRes := userService.PostUser(paramUser)

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
