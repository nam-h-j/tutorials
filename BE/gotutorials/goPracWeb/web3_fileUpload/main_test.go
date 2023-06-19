package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUploadTest(t *testing.T) {
	// assert를 불러옴
	assert := assert.New(t)
	// 업로드할 파일을 확인
	path := "~/Desktop/testimg/lion.png"
	file, _ := os.Open(path) // 파일을 열어본다.
	defer file.Close()       //os의 파일 자원을 닫는다

	// 업로드할 폴더 내용을 다 지움
	os.RemoveAll("./upload")

	// 버퍼에 임시저장된 파일을 체크하기
	buf := &bytes.Buffer{}
	// 파일 전송하는 포맷인 MIME에 버퍼를 전달
	writer := multipart.NewWriter(buf)
	// 폼에서 받은 파일을 생성
	multi, err := writer.CreateFormFile("upload_file", filepath.Base(path))
	assert.NoError(err)  // 파일 생성시 에러 체크
	io.Copy(multi, file) // 폼 파일 카피
	writer.Close()       // 파일 카피 했으므로 종료

	// 테스트 요청 응답 작성
	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/uploads", buf)
	req.Header.Set("Content-type", writer.FormDataContentType()) // 데이터 타입 정의

	// 함수에 테스트 요청 보내기
	uploadHandler(res, req)
	assert.Equal(http.StatusOK, res.Code)

	// 위치에 있는 파일이 제대로 올라갔는지 확인
	uploadFilePath := "./upload/" + filepath.Base(path)
	_, err = os.Stat(uploadFilePath) // 파일의 정보를 가져다줌, (리턴밸류가 필요없으면 명명하지 않고 "_"를 쓰면 됨)
	assert.NoError(err)

	uploadFile, _ := os.Open(uploadFilePath) // 업로드한 파일을 오픈
	originFile, _ := os.Open(path)           // 원본 파일을 오픈
	defer uploadFile.Close()
	defer originFile.Close()

	// 데이터를 비교를 위해서 바이트 데이터를 읽음
	uploadData := []byte{}
	originData := []byte{}
	uploadFile.Read(uploadData)
	originFile.Read(originData)

	// 두 바이트 값을 비교
	assert.Equal(originData, uploadData)
}
