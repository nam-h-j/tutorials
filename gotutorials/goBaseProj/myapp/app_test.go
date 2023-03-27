package myapp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler()) // 목업 서버 열기
	defer ts.Close()                       // 목업 서버 닫기

	// 요청테스트
	res, err := http.Get(ts.URL)
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)

	// 출력테스트
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("this is index", string(data))
}

func TestUser(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	res, err := http.Get(ts.URL + "/user")
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)

	data, _ := ioutil.ReadAll(res.Body)
	assert.Contains(string(data), "user, get user info : ")
}

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

func TestPutUser(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	// 요청 테스트시 PUT역시 Go에서 기본으로 제공되지 않으므로
	//NewRequest + DefaultClient.Do로 요청해야함
	req, _ := http.NewRequest("PUT", ts.URL+"/user",
		strings.NewReader(`{"id":1, "first_name":"updated", "last_name":"updated", "email":"updated@naver.com"}`))
	res, err := http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Contains(string(data), "No User ID : 1")

	// post 하고나서 update 테스트 해보기
	// 1. post 요청하기
	res, err = http.Post(ts.URL+"/user", "application/json",
		strings.NewReader(`{"first_name":"jackson", "last_name":"nam", "email":"jackson@naver.com"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, res.StatusCode)

	//2. post 한 유저 정보 ID 뽑아내기
	user := new(User)
	err = json.NewDecoder(res.Body).Decode(user)
	assert.NoError(err)
	assert.NotEqual(0, user.ID)

	//3. 업데이트하기
	updateStr := fmt.Sprintf(`{"id":%d, "first_name":"updated"}`, user.ID)
	req, _ = http.NewRequest("PUT", ts.URL+"/user",
		strings.NewReader(updateStr))
	res, err = http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)

	//4. 업데이트 제대로 되었는지 확인
	updateUser := new(User)
	err = json.NewDecoder(res.Body).Decode(updateUser)
	assert.NoError(err)
	assert.Equal(updateUser.ID, user.ID)             //아이디가 같은지 검사
	assert.Equal("updated", updateUser.FirstName)    //데이터가 변경되었는지 검사
	assert.Equal(user.LastName, updateUser.LastName) //데이터가 그대로인지 검사
	assert.Equal(user.Email, updateUser.Email)       //데이터가 그대로인지 검사
}
