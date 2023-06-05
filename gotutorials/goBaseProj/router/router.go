package router

import (
	"database/sql"

	"../handler/user"
	"../middleware"

	"github.com/gin-gonic/gin"
)

func Router(db *sql.DB) *gin.Engine {
	// userMap = make(map[int]*User) //init userMap
	// lastID = 0
	router := gin.Default()

	// 라우터에서 디비 미들웨어를 통해 디비 풀을 재사용 하도록 함
	router.Use(middleware.MiddleDB(db))

	// index
	// router.GET("/", indexHandler)

	userRouter := router.Group("user")
	userRouter.GET("/", user.GetUserList)
	userRouter.GET("/:id", user.GetUser)
	userRouter.POST("/", user.PostUser)
	userRouter.DELETE("/:id", user.DeleteUser)
	userRouter.PUT("/", user.PutUser)
	// router.PUT("/user", updateUserHandler)

	return router
}
