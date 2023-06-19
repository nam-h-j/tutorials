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
- Prepared Statements를 사용한 쿼리 요청

## Prepared Statements

- SQL 쿼리를 미리 컴파일하고 나중에 실행할 때 매개변수를 전달하는 방법
- 동적으로 변경되는 매개변수를 쿼리에 전달 할 수 있으므로 SQL 인젝션 공격을 방지할 수 있는 장점이 있다.
- 적용 예

```go
// Prepared Statements 작성
deleteQuery := "DELETE FROM some_data_base WHERE column1 = ? AND column2 = ?"
deleteStmt, err := self.DB.Prepare(deleteQuery)
if err != nil {
	fmt.Println(err)
	continue
}
defer deleteStmt.Close()

// 매개변수 전달 및 쿼리 실행
if _, err := deleteStmt.Exec(param1, param2); err != nil {
	fmt.Println(err)
	continue
}
```

## 추가로 진행할 작업

- 로그인(JWT), 인증 회원가입
-
