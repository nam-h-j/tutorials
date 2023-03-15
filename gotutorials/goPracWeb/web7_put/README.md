# web5 post와 post테스트

## 소스코드내용
### 1_NewHandler
- userMap과 lastID를 초기화
- getUserHandler를 호출
```go
func NewHandler() http.Handler {
	userMap = make(map[int]*User) //init userMap
	lastID = 0 //user Idx
	mux := mux.NewRouter()

    ...
	mux.HandleFunc("/user/{id:[0-9]+}", getUserHandler).Methods("GET")
	...

    return mux
}
```

### 2_getUserHandler
- getUserHandler의 요청파라미터인 id를 취득
- 스트링 형태인 id를 int로 변환
- 만약 요청파라이터에 id가 없다면, 또는 잘못된 형식이라면 StatusBadRequest와 err 내용을 출력하고 리턴
- user Map에 요청한 id가 있는지 체크
- 만약 요청한 id가 없다면 No User ID : ${id} 메시지 출력
- 만약 요청한 id가 있으면 해당 유저의 정보를 출력한다.
```go
func getUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	user, ok := userMap[id]
	if !ok {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No User ID : ", id)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(user)
	fmt.Fprint(w, string(data))
}
```
***

## 테스트코드
### 1_TestGetUser
- 존재하지 않는 유저 idx를 요청
- No User ID : ${idx}를 출력
```go
func TestGetUser(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	res, err := http.Get(ts.URL + "/user/89")
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Contains(string(data), "No User ID : 89")
}
```
### 2_TestPostUser
- 새로운 유저를 post 요청
- post요청의 response 된 유저 id 값이 0이 아닌지 확인
- post한 유저 정보가 정상적으로 get되는지 확인
    - post 요청에서 받은 user.ID 값으로 get요청을 하기 위해서 int에서 string으로 캐스팅
    - get 요청이 정상인지 확인
    - get 으로 넘어온 데이터가 post resp 값과 같은지 비교

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
