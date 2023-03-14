# web1 handler

## 소스코드내용
### 1_NewHandler
```go
func main() {
	http.ListenAndServe(":1234", myapp.NewHttpHandler())
}
```