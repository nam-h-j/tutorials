package router

import (
	"database/sql"

	"goBaseProj/docs"
	"goBaseProj/handler/user"
	"goBaseProj/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router(db *sql.DB) *gin.Engine {
	// userMap = make(map[int]*User) //init userMap
	// lastID = 0
	router := gin.Default()

	docs.SwaggerInfo.BasePath = "/"

	gin.SetMode(gin.ReleaseMode)

	// 라우터에서 디비 미들웨어를 통해 디비 풀을 재사용 하도록 함
	router.Use(middleware.MiddleDB(db))

	// router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// index
	// router.GET("/", indexHandler)
	router.GET("/welcome/:name", welcomePathParam)

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	userRouter := router.Group("user")
	userRouter.GET("/", user.GetUserList)
	userRouter.GET("/:id", user.GetUser)
	userRouter.POST("/", user.PostUser)
	userRouter.DELETE("/:id", user.DeleteUser)
	userRouter.PUT("/", user.PutUser)
	// router.PUT("/user", updateUserHandler)

	return router
}

type welcomeModel struct {
	ID   int    `json:"id" example:"1" format:"int64"`
	Name string `json:"name" example:"account name"`
}

// Welcome godoc
// @Tags Welcome!
// @Summary 스웨거 테스트용 핸들러
// @Description 스웨거 테스트용 핸들러 Desc
// @name get-string-by-int
// @Accept  json
// @Produce  json
// @Param name path string true "User name"
// @Router /welcome/{name} [get]
// @Success 200 {object} welcomeModel
func welcomePathParam(c *gin.Context) {
	name := c.Param("name")
	message := name + " is very handsome!"
	welcomeMessage := welcomeModel{1, message}
	c.JSON(200, gin.H{"message": welcomeMessage})
}
