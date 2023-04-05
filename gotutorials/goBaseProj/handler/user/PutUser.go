package user

import (
	"database/sql"
	"net/http"

	"../../model"
	"./service"

	"github.com/gin-gonic/gin"
)

func PutUser(c *gin.Context){
	resBody := model.User{}

	// 디비풀 연결
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		panic(ok)
	}

	// response 객체 검사
	err := c.BindJSON(&resBody)
	if err != nil{
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