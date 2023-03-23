# web5_ㅔㅐㄴㅅ

## SOURCE CODE
### NewHandler
- 요청 핸들러 등록
#### GIN
```go
func NewHandler() *gin.Engine {
	router := gin.Default()
	...
	router.POST("/user", createUserHandler)
	...
	return router
}
```
#### MUX
```go
func NewHandler() http.Handler {
	mux := mux.NewRouter()
	...
	mux.HandleFunc("/user", createUserHandler).Methods("POST")
	...
	return mux
}
```
### createUserHandler
#### GIN
- 유저객체를 생성
- body에 넘어온 json 스트링을 디코드
- 값이 없으면 에러처리
- 객체에 값 할당
- 값이 할당된 객체를 마샬링하고 출력
```go
func createUserHandler(c *gin.Context) {
	user := new(User) // create User struct
	err := json.NewDecoder(c.Request.Body).Decode(user)
	if err != nil {
		c.String(http.StatusOK, err.Error())
		return
	}

	// Created User
	lastID++
	user.ID = lastID
	user.CreatedAt = time.Now()
	userMap[user.ID] = user

	c.Writer.Header().Set("Content-Type", "application/json")
	data, _ := json.Marshal(user) // 마샬링, 논리적 구조를 로우바이트로 변경하는 것(인코딩)
	c.String(http.StatusCreated, string(data))
}
```
#### MUX
```go
func createUserHandler(w http.ResponseWriter, r *http.Request) {
	user := new(User) // create User struct
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	// Created User
	lastID++
	user.ID = lastID
	user.CreatedAt = time.Now()
	userMap[user.ID] = user

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	data, _ := json.Marshal(user) // 마샬링, 논리적 구조를 로우바이트로 변경하는 것(인코딩)
	fmt.Fprint(w, string(data))
}
```
## TEST CODE
### TestPostUser
- POST 요청을 보냄, 바디에 테스트용 json 스트링을 같이 보냄
- POST가 성공적으로 진행되었는지 확인
- POST된 데이터 확인을 위해 POST한 id로 GET요청 보냄
- GET 요청 바디로 넘어온 데이터가 맞는지 확인
```go
func TestPostUser(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	//1. 요청테스트
	res, err := http.Post(ts.URL+"/user", "application/json", strings.NewReader(`{"first_name":"jackson", "last_name":"nam", "email":"now@naver.com"}`))

	assert.NoError(err)
	assert.Equal(http.StatusCreated, res.StatusCode)

	//2. 데이터 테스트 유저정보 post 되었는지 테스트
	user := new(User)                            // 유저객체 생성
	err = json.NewDecoder(res.Body).Decode(user) //리스폰스 바디를 json으로 디코딩
	assert.NoError(err)
	assert.NotEqual(0, user.ID) //유저의 ID가 0이 아니어야 함.

	// 3. post 한 유저 정보로 get 되는지 테스트
	id := user.ID
	res, err = http.Get(ts.URL + "/user/" + strconv.Itoa(id)) //strconv는 id가 int이기 때문에 url에서 사용하기 위해서 str로 캐스팅.
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode) //요청정상인지 확인

	// 4. 바디 데이터 값 정상인지 확인
	userBodyCompair := new(User) //비교 하기 위한 user 객체 생성
	err = json.NewDecoder(res.Body).Decode(userBodyCompair)
	assert.NoError(err)
	assert.Equal(user.ID, userBodyCompair.ID)
	assert.Equal(user.FirstName, userBodyCompair.FirstName)

}
```