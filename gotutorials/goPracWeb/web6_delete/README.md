# web6 Delete

## 소스코드
### 1_NewHandler
- deleteUserHandler를 호출
```go
func NewHandler() http.Handler {
    ...
	mux.HandleFunc("/user/{id:[0-9]+}", deleteUserHandler).Methods("DELETE")
	...
    return mux
}
```

### 2_deleteUserHandler
- deleteUserHandler의 요청파라미터인 id를 취득
- 스트링 형태인 id를 int로 변환
- 만약 요청파라이터에 id가 없다면, 또는 잘못된 형식이라면 StatusBadRequest와 err 내용을 출력하고 리턴
- user Map에 요청한 id가 있는지 체크
- 만약 요청한 id가 없다면 No User ID : ${id} 메시지 출력
- 만약 요청한 id가 있으면 해당 유저의 정보를 삭제한다.
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
***

## 테스트코드
### 1_TestDeleteUser
- 요청 테스트시 Delete는 Go에서 기본으로 제공되지 않으므로, NewRequest + DefaultClient.Do로 요청
- response body에 해당 id의 객체가 없다면 No User ID : {id}를 출력하는지 검사
- post 한 다음 유저가 생성, 삭제 되는지 테스트

```go
func TestDeleteUser(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	// 요청 테스트시 Delete는 Go에서 기본으로 제공되지 않으므로
	// NewRequest + DefaultClient.Do로 요청해야함
	req, _ := http.NewRequest("DELETE", ts.URL+"/user/1", nil)
	res, err := http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Contains(string(data), "No User ID : 1")

	//post 한 다음 유저가 생성, 삭제 되는지 테스트
	//1. Post 요청
	res, err = http.Post(ts.URL+"/user", "application/json",
		strings.NewReader(`{"first_name":"jackson", "last_name":"nam", "email":"now@naver.com"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, res.StatusCode)

	//2. 데이터 테스트 유저정보 post 되었는지 테스트
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
