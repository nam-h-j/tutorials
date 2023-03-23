# web4_GET

## SOURCE CODE
### NewHandler
- 요청 핸들러 등록
#### GIN
```go
func NewHandler() *gin.Engine {
	router := gin.Default()
	...
	router.GET("/user/:id", getUserHandler)
	...
	return router
}
```
#### MUX
```go
func NewHandler() http.Handler {
	mux := mux.NewRouter()
	...
	mux.HandleFunc("/user/{id:[0-9]+}", getUserHandler)
	...
	return mux
}
```
### getUserHandler
- get요청의 id param을 받아서 "user id : id" 형식으로 출력
#### GIN
```go
func getUserHandler(c *gin.Context) {
	id, _ := c.Params.Get("id")
	c.String(http.StatusOK, "user id: %s", id)
}
```
#### MUX
```go
func getUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, "user id: ", vars["id"])
}
```
## TEST CODE
### TestGetUser
- 89, 100 두가지 id를 담아서 GET 요청이 정상인지 확인하는 테스트 코드
```go
func TestGetUser(t *testing.T) {
	assert := assert.New(t)

	// 테스트서버를 띄운다.
	ts := httptest.NewServer(NewHandler())
	// 테스트서버를 띄우고나서 닫아준다.
	// defer : 이 함수가 종료되기 직전(리턴직전)에 실행한다
	defer ts.Close()

	res, err := http.Get(ts.URL + "/user/89")
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)

	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal(string(data), "user id: 89")

	res, err = http.Get(ts.URL + "/user/100")
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)

	data, _ = ioutil.ReadAll(res.Body)
	assert.Equal(string(data), "user id: 100")
}
```