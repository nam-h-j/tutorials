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
- /user 요청이 왔을때 barHandler를 실행한다.
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
### 3_barHandler
- user객체가 response로 넘어오면 Decode 한다
- 
```go
func barHandler(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Request: ", err)
		return
	}
	user.CreatedAt = time.Now()
	data, _ := json.Marshal(user)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, string(data))
}
```
