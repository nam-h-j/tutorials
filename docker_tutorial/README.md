## 도커 기초
### 작성시 유의 사항
- 레이어 순서대로 작성한다.(캐싱되기 때문에 효율이 높아짐)
- 베이스이미지,
- 컨테이너 세팅(실행환경)
- 실행

### DockerFile 설명
```dockerfile
# FROM baseImage를 항상 로드 해야함,
# node:16-alpine <= 노드에서 미리 제공하는 노드전용 도커 이미지
# node를 ctrl + click 하면 노드 도커 이미지 관련 문서 볼 수 있음
FROM node:16-alpine 

# 작업할 도커 디렉토리 선택, (WORKDIR = 리눅스의 CD와 같음)
WORKDIR /app

# 노드 패키지 정보 및 소스파일 카피
COPY package.json package-lock.json ./

# 컨테이너에 노드 환경 설치
# RUN npm install = package.json 정보에 있는 노드에 호환가능한 버전을 설치(다른 버전으로 설치 될 수도 있음)
# RUN npm ci = package-lock.json 정보에 있는 노드를 버전을 그대로 설치(버전이 안바뀜)
RUN npm ci

# 실행파일 복사
COPY index.js .

# 실행하기 노드로 index.js를 실행하기
ENTRYPOINT [ "node", "index.js" ]
```

### 도커 기초 사용법

#### 도커 이미지 생성하기
```
$ docker build -f Dockerfile -t test-docker-image .
```
- -f Dockerfile : 어떤 도커 파일을 사용할지 설정(-f name of the dockerfile)
- -t fun-docker : 도커 이미지 이름을 설정
- . : 도커파일이 있는 위치

#### 도커 이미지 확인하기
```
$ docker images
```
#### 도커 실행하기
```
$ docker run -d -p 8080:8080 test-docker-image
```
- -d : stands for detached : 컨테이너 백그라운드에서 실행
- -p 8080:8080 : 호스트 포트와 도커 포트 연결 (hostport:dockerport)

#### 실행중인 컨테이너 확인하기
```
$ docker ps
```

#### 컨테이너 정보 확인하기
```
$ docker logs {CONTAINER ID}
```

#### 컨테이너 레지스트리에 푸시하기(도커허브 dockerhub)
1. 도커허브에서 리포지토리 생성
2. 리포지토리 생성하면 리포지토리 주소가 생성되는데 (userid/repo-name)이 이름과 이미지 파일 이름을 매칭시켜야함
```
# 이미지 이름 추가하는 방법
$ docker tag test-docker-image:latest userid/test-docker-image:latest
$ docker tag {도커이미지이름:tag} {변경할이름:tag}
```
3. docker images 명령어로 이름이 들어가있는지 확인
4. 로그인
```
$ docker login
```
5. 푸시
```
$ docker push userid/test-docker-image:latest
```
