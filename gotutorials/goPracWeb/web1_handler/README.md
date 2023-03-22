# web1_handler

## 소스코드

### 1_NewHandler
- myapp 패키지의 NewHttpHandler를 호출, 서버실행
#### GIN
```go
func main() {
	router := myapp.NewHttpHandler()
	router.Run(":1234")
}
```
#### MUX
```go
func main() {
	http.ListenAndServe(":1234", myapp.NewHttpHandler())
}
```
### 2_NewHttpHandler
- root('/') 요청이 왔을때 "myappRootOk"를 페이지에 출력한다.
- /user 요청이 왔을때 barHandler를 실행한다.
#### GIN
```go
func NewHttpHandler() *gin.Engine {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "myappRootOk")
	})
	router.GET("/user", userHandler)
	return router
}
```
#### MUX
- root('/') 요청이 왔을때 "myappRootOk"를 페이지에 출력한다.
- /user 요청이 왔을때 barHandler를 실행한다.
```go
func NewHttpHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "myappRootOk")
	})

	mux.HandleFunc("/user", userHandler)
	return mux
}
```
### 3_userHandler
- user객체가 response로 넘어오면 go객체로 언마샬링 한다
	- 마샬링 : GO -> JSON
	- 언마샬링 : JSON -> GO
- 유저 정보가 없으면 BadReq를 출력
- 유저 정보가 있으면 마샬링후 화면에 데이터 프린트
#### GIN
```go
func userHandler(c *gin.Context) {
	user := new(User)
	err := c.BindJSON(&user)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad Request: ", err.Error())
		return
	}
	user.CreatedAt = time.Now()
	data, _ := json.Marshal(user)
	c.IndentedJSON(http.StatusCreated, data)
}
```
#### MUX
```go
func barHandler(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Request: ", err)
		return
	}
	user.CreatedAt = time.Now()
	data, _ := json.Marshal(user)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, string(data))
}
```
