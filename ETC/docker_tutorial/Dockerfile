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
