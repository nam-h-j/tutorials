# Javascirpt postMessage() test

### 아이프레임 -> 부모html 값 넘겨주기

> 아이프레임 쪽 코드 데이터를 보내는 부분</br>
> data : 아이프레임에서 부모 html에 넘길 값</br>
> origin : host를 입력 같은 host가 아니면 cors</br>
```
//setter
window.addEventListener("load", function () {
  window.parent.postMessage(data, origin);
});
//setter
```

</br>

> 부모 html쪽에서 데이터를 받는 부분<br>
> message 이벤트를 걸어두면 아이프레임 쪽에서 postMessage가 실행되었을때 e.data로 값이 넘겨진다<br>
```
//getter
window.addEventListener("message", function (e) {
    console.log(e.data);
});
//getter
```

