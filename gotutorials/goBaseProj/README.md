# GO_BASE_PROJECT
- gin web framework 사용
## 구조
- main -> middleWare -> router -> handler -> service

## middleWare
- 프로그램 시작전에 미리 실행 하는 기능들을 정의.
- database.go => DB커넥션 풀을 생성하는 middleware
```go
package middleware

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func MiddleDB(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("databaseConn", db)
		c.Next()
	}
}
```
### DB connection
## service
- DB 쿼리가 담기는 부분
## 추가로 진행할 작업
- 로그인(JWT), 인증 회원가입
- 
