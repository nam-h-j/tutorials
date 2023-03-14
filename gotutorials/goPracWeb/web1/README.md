# web1 handler

## 소스코드

### 1_NewHandler
- myapp 패키지의 NewHttpHandler를 호출
```go
func main() {
	http.ListenAndServe(":1234", myapp.NewHttpHandler())
}
```
### 2_NewHttpHandler
- root('/') 요청이 왔을때 "myappRootOk"를 페이지에 출력한다.

```go
func NewHttpHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "myappRootOk")
	})

	mux.HandleFunc("/user", barHandler)
	return mux
}
```