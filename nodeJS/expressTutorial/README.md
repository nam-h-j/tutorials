## express 시작하기

### 설치
```
npm init -y
npm install express
```

### 사용한 패키지
- bodyParser
    - bodyParser, 미들웨어, 요청과 응답을 조작할 수 있게 해줌
    - request의 body를 원하는 형태로 파싱해줌
- nodemon(dev로 설치 --save-dev)
    - 소스 변경시 node 서버를 자동으로 재시작해주는 툴
- uuid
    - 네트워크 상에 고유성이 보장되는 id를 만들기 위한 표준 규약
    - UUID는 Universally Unique IDentifier의 약어이고 범용 고유 식별자라고 한다.
    - UUID는 128비트의 숫자이며, 32자리의 16진수로 표현된다.
    - 8자리-4자리-4자리-4자리-12자리 패턴으로 하이픈을 집어 넣어 5개의 그룹으로 구분한다.
    - [uuid 자세히..](https://mattmk.tistory.com/31)
### package.json 설정
- module 설정
```js
...
"main": "index.js",
"type": "module", // <= 이 옵션을 추가해주면 import "" from "" 키워드로 패키지 불러오기 가능
...
```
- nodemon 설정
```js
"scripts": {
    "start": "nodemon index.js"
},
```

