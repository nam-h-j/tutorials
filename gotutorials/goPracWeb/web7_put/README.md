# web7_put

## SOURCE CODE
### NewHandler
- 요청 핸들러 등록
#### GIN
```go
func NewHandler() *gin.Engine {
	router := gin.Default()
	...
	router.DELETE("/user/:id", deleteUserHandler)
	...
	return router
}
```
#### MUX
```go
func NewHandler() http.Handler {
	mux := mux.NewRouter()
	...
	mux.HandleFunc("/user/{id:[0-9]+}", deleteUserHandler).Methods("DELETE")
	...
	return mux
}
```
### deleteUserHandler
- 아이디 파라미터 받기, int로 캐스팅
- 유저 맵에 해당 id가 있는지 검사
- 없으면 No User ID 출력
- 있으면 삭제하고 Deleted User ID 출력
#### GIN
```go
func deleteUserHandler(c *gin.Context) {
	idStr, _ := c.Params.Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	_, ok := userMap[id]
	if !ok {
		c.String(http.StatusOK, "No User ID : %v", id)
		return
	}
	delete(userMap, id) //해당 키를 가진 객체를 Map에서 삭제한다
	c.String(http.StatusOK, "Deleted User ID : %v", id)
}
```
#### MUX
```go
func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	_, ok := userMap[id]
	if !ok {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No User ID : ", id)
		return
	}
	delete(userMap, id) //해당 키를 가진 객체를 Map에서 삭제한다
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted User ID : ", id)
}
```
## TEST CODE
### TestDeleteUser
- 요청 테스트시 Delete는 Go에서 기본으로 제공되지 않으므로
- NewRequest + DefaultClient.Do로 요청해야함
- 삭제할 값이 없는 경우를 테스트
	- Delete 요청을 보냄 삭제할 id 값을 보냄
	- 요청 확인
	- 출력 확인
- 삭제할 값이 있는 경우를 테스트
	- 먼저 USER 생성을 위해서 POST요청 후 검사
	- POST한 USER가 등록이 되었는지 데이터 검사
	- DELETE 요청을 검사
```go
func TestDeleteUser(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	// 요청 테스트시 Delete는 Go에서 기본으로 제공되지 않으므로
	//NewRequest + DefaultClient.Do로 요청해야함
	req, _ := http.NewRequest("DELETE", ts.URL+"/user/1", nil)
	res, err := http.DefaultClient.Do(req)
	
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)
	
	data, _ := ioutil.ReadAll(res.Body)
	assert.Contains(string(data), "No User ID : 1")

	//등록해서 삭제 되는지 확인해보기
	//1. Post 요청 검사
	res, err = http.Post(ts.URL+"/user", "application/json",
		strings.NewReader(`{"first_name":"jackson", "last_name":"nam", "email":"now@naver.com"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, res.StatusCode)

	//2. post 되었는지 userid 검사
	user := new(User)
	err = json.NewDecoder(res.Body).Decode(user)
	assert.NoError(err)
	assert.NotEqual(0, user.ID)

	//3. 삭제 요청 다시 날려서 삭제 확인
	req, _ = http.NewRequest("DELETE", ts.URL+"/user/1", nil)
	res, err = http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)
	
	data, _ = ioutil.ReadAll(res.Body)
	assert.Contains(string(data), "Deleted User ID : 1")
}
```