# DynamicForms_Vuetify
- Nuxt.js, Vuetify를 활용한 폼 생성 컴포넌트 데모.

## 초기 셋업
### 폰트적용(Noto Sans KR)
#### 웹폰트 불러오기
- nuxt.config.js
```js
head: {
    ...
    link: [... ,{rel: 'stylesheet', href: 'https://fonts.googleapis.com/css2?family=Noto+Sans+KR:wght@300;400;500;700&display=swap'}],
  },
```

#### scss 글로벌 스타일 생성 및 웹 폰트 적용
- assets/scss/Global.scss
```css
* {
  font-family: 'Noto Sans KR', 'Roboto', 'sans-serif';
}
```

#### 글로벌 scss 파일 설정파일에 적용
- nuxt.config.js
```js
  ...
  css: ['~/assets/scss/Global.scss'],
  ...
```

## 페이지 및 기능 구성
- 입력모듈관리하기
  - 모듈 생성하기
    - 모듈 생성 페이지
      > 작성 및 수정
        > 텍스트, 숫자, 전화번호, 주소, 체크박스, 라디오버튼, 셀렉트박스, 파일, 이미지
      > 기능연결
        > 업데이트시 알림, 회계기능 등
      > 미리보기
      > 저장
  - 생성된 모듈의 목록
    > 미리보기
    > 수정
    > 삭제
