# JWT with Go(with gin)

## Library

```go
import (
	...
  "github.com/go-redis/redis/v7" // redis driver
	"github.com/dgrijalva/jwt-go" // jwt token generator
	"github.com/gin-gonic/gin" // gin
)
```

## Router

```go
var (
  router = gin.Default()
)

func main() {
  router.POST("/login", Login)
  router.Run(":8080")
}
```

## User Struct

```go
type User struct {
  ID       uint64    `json:"id"`
  Username string    `json:"username"`
  Password string    `json:"password"`
  Phone    string    `json:"phone"`
}
```

## Mock Data

```go
var user = User{
  ID:            1,
  Username: "username",
  Password: "password",
  Phone:    "49123454322",
}
```

## Login

```go
func Login(c *gin.Context) {
  // 1. 요청을 사용자 구조체에 언마샬링
  var u User
  if err := c.ShouldBindJSON(&u); err != nil {
     c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
     return
  }

  // 2. Mock 사용자와 비교
  if user.Username != u.Username || user.Password != u.Password {
     c.JSON(http.StatusUnauthorized, "Please provide valid login details")
     return
  }

  // 3. user.ID를 넘겨서 토큰을 생성
  token, err := CreateToken(user.ID)
  if err != nil {
     c.JSON(http.StatusUnprocessableEntity, err.Error())
     return
  }
  // 4. 토큰을 반환
  c.JSON(http.StatusOK, token)
}
```

## CreateToken

```go
// Create Token
func CreateToken(userId uint64) (string, error) {
  var err error
  //Creating Access Token
  os.Setenv("ACCESS_SECRET", "jdnfksdmfksd")
  atClaims := jwt.MapClaims{}
  atClaims["authorized"] = true
  atClaims["user_id"] = userId
  // 토큰 유효 시간
  atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
  at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
  token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
  if err != nil {
     return "", err
  }
  return token, nil
}
```
