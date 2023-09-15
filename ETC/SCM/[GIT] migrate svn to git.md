# migrate svn to git (with all svn commit histories )
- using 'git-svn' then it works so easily
  - git-svn을 활용하면 아주 손쉽게 가능함
## git-svn 설치
```bash
$ brew install git-svn
```

## 빈폴더 생성
```bash
$ mkdir folder
$ cd folder
```

## git-svn 명령어로 git-repository 초기화 및 svn repo로 부터 데이터 가져오기
```bash
$ git svn init [svn repository url] -T trunk -b braches -t tags
```
- svn 전체가 아닌 일부 폴더만 이동 하려면 해당 패스를 작성해주어야 한다.
- trunk, branches, tags이외에 특정 이름으로 세팅이 되어있다면 설정이 필요하다.

## svn 히스토리 가져오기
```bash
$ git svn fetch
```
- 중단시 다시 실행하면 중단된 시점부터 다시 이력을 가지고 온다.
- 변경사항이 있다면 해당 명령어로 최신화 시켜줄수 있다.

## 로컬, 리모트 브랜치 확인
```bash
& git branch -a
```
- svn의 브랜치 정보가 git으로 제대로 옮겨져 왔는지 확인

## push할 git 저장소 remote url을 설정
```bash
$ git remote add origin [git repository url]
& git remote -v
```

## 저장소로 push
```bash
$ git push origin master
```
