# web8 dbConnection with mariaDb

## 작업순서
### 마리아DB 설치
- windows : 마리아db 공식 사이트에서 다운로드 및 설치
- mac : homebrew로 설치
```
- 마리아디비 설치
$ brew install mariadb

- 마리아디비 시작 
$ mysql.server start
또는
$ brew services start mariadb
```
### 마리아DB GO 드라이버 설치
```
$ go get github.com/go-sql-driver/mysql
```
### os 환경변수 설정
- db 접속에 필요한 계정 정보를 os 환경변수에 등록해둔다.
- windows
```
C:\Users\you\data-access> set DBUSER=username
C:\Users\you\data-access> set DBPASS=password
```
- mac
```
$ export DBUSER=username
$ export DBPASS=password
// 환경변수 확인
$ printenv
```
### main.go 에서 DB 연결하기
- Connected to mysql version: "version" 이 출력되면 연결 성공
```go
// DB Config
cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "go_crud",
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Println("Connected to mysql version: ", version)

	http.ListenAndServe(":1234", myapp.NewHandler())
```