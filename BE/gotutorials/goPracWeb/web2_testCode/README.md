# web2_testCode

## SOURCE CODE
### indexHandler
- 인덱스 페이지에 "myapp_index" 출력
#### GIN
```go
func indexHandler(c *gin.Context) {
	c.String(http.StatusOK, "myapp_index")
}
```
#### MUX
```go
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "myapp_index")
}
```
## TEST CODE
### TestIndexHandler
- "github.com/stretchr/testify/assert" 패키지 사용
- assert 초기화
- res := httptest.NewRecorder()에 req 테스트 결과가 들어감
- req := httptest.NewRequest("", "/", nil) 테스트할 api 사양 작성
	- 첫번째 파라미터 = 테스트 요청 방식, 공백으로 두면 GET
	- 두번째 파라미터 = URL
	- 세번째 파라미터 = Header에 보낼 값, 없으면 nil
- 테스트할 router를 불러오기
- res, req로 서버에 요청을 날린다.
- http통신 STATUS 체크
- res에 담긴 내용을 읽어온다.
- Body 내용이 맞는지 체크
#### GIN
```go
func TestIndexHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("", "/", nil)

	router := NewHttpHandler()
	router.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)

	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("myapp_index", string(data))
}
```
#### MUX
```go
func TestIndexHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("", "/", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
 
 	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("myapp_index", string(data))
}
```
## SOURCE CODE
### getUserNameFromParam
- URL 쿼리에서 유저이름을 취득해서 출력하는 로직
#### GIN
```go
func getUserNameFromParam(c *gin.Context) {
	name := c.DefaultQuery("name", "Guest")
	c.String(http.StatusOK, "welcome %s!", name)
}
```
- gin에서 쿼리를 취득하는 3가지 방법
```go
//1. 
name := c.Request.URL.Query().Get("name")
//2. 1번 의 shortCut
name := c.Query("name")
//3. 디폴트 밸류 설정 with default values
name := c.DefaultQuery("name", "Guest")
```
#### MUX
```go
func getUserNameFromParam(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "unknown"
	}
	fmt.Fprintf(w, "welcome %s!", name)
}
```
## TEST CODE
### TestGetUserNameFromParam_WithoutParam
- 파라미터가 없는 경우를 테스트
```go
func TestGetUserNameFromParam_WithoutParam(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/user", nil)

	router := NewHttpHandler()
	router.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("welcome Guest!", string(data))
}
```
### TestGetUserNameFromParam_WithoutParam
- 파라미터가 있는 경우를 테스트
- 더미 텍스트를 만들어서 테스트 해본다
```go
func TestGetUserNameFromParam_WithParam(t *testing.T) {
	assert := assert.New(t)

	mockName := "Randy"
	param := fmt.Sprintf("/user?name=%s", mockName)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", param, nil)

	router := NewHttpHandler()
	router.ServeHTTP(res, req)
	assert.Equal(http.StatusOK, res.Code)

	data, _ := ioutil.ReadAll(res.Body)
	resultStr := fmt.Sprintf("welcome %s!", mockName)
	assert.Equal(resultStr, string(data))
}
```