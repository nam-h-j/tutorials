package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	// input에서 업로드한 파일 받기
	uploadFile, header, err := r.FormFile("upload_file")
	// 에러처리
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	// 파일을 받은뒤에 파일 업로드 자원을 닫기
	defer uploadFile.Close()

	// 저장할 디렉터리 이름 및 패스 지정
	dirname := "./upload"
	// 디렉터리가 없으면 생성
	os.MkdirAll(dirname, 0777)
	// 저장될 파일 패스와 파일 명을 지정
	filepath := fmt.Sprintf("%s/%s", dirname, header.Filename)
	// 파일을 생성
	file, err := os.Create(filepath)
	// 파일 생성후 파일 생성 자원 닫기
	defer file.Close()
	// 에러처리
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
	}
	// 생성된 파일 자원에 upload된 파일을 복사
	io.Copy(file, uploadFile)
	// 성공처리
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, filepath)

}

func main() {
	http.HandleFunc("/upload", uploadHandler)
	http.Handle("/", http.FileServer(http.Dir("public")))

	http.ListenAndServe(":1234", nil)
}
