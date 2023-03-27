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
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// 데이터베이스 풀 커넥션 재사용
	router.Use(middleware.MiddleDB(db))

	// index
	// router.GET("/", indexHandler)

	userRouter := router.Group("user")
	userRouter.POST("/post", user.PostUser)

	// userRouter.GET("/user", user.PostUser)
	// router.GET("/user/:id", getUserHandler)
	// router.POST("/user", createUserHandler)
	// router.PUT("/user", updateUserHandler)
	// router.DELETE("/user/:id", deleteUserHandler)

	return router
}
