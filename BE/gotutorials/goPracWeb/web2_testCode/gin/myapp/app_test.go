package myapp

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	//assert : 테스트 한 값이 맞는지 비교하는데 사용
	//이 패키지를 사용하지 않으면 일일이 테스트 결과에 대한 분기 작성해줘야함
	"github.com/stretchr/testify/assert"
)

func TestIndexHandler(t *testing.T) {
	//assert 생성
	assert := assert.New(t)
	//테스트할 핸들러가 있는 라우터를 호출
	router := NewHttpHandler()

	//1. 테스트 할 req와 res를 작성해준다.
	res := httptest.NewRecorder()
	req := httptest.NewRequest("", "/", nil)

	//2. ginServer에 res와 req를 넘겨준다.
	router.ServeHTTP(res, req)

	//3. 테스트 결과 보기
	//3-1 요청이 제대로 처리되었는지 확인
	assert.Equal(http.StatusOK, res.Code)
	//3-2 바디에 원하는 결과가 제대로 나왔는지 확인
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("myapp_index", string(data))
}

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
