window.onload = function(){
//DOM을 가지고 오는 상수
  const clockWrap = document.querySelector(".clock"),
        clock = clockWrap.querySelector("h1"),
        form = document.querySelector(".form"),
        input = form.querySelector("input"),
        greetings = document.querySelector(".greetings"),
        toDoForm = document.querySelector(".toDoForm"),
        toDoInput = toDoForm.querySelector("input"),
        toDoList = document.querySelector(".toDoList")

//함수에 필요한 상수값들
  const userStorage = "currentUser",
        visible = "visible",
        toDoStorage = "toDoItem",
        toDoArray = [];

//현재시간을 가지고 오는 함수
  function getTime(){
    const date = new Date();
    const minutes = date.getMinutes() < 10 ? "0" + date.getMinutes() : date.getMinutes();
    const hours = date.getHours() < 10 ? "0" + date.getHours() : date.getHours();
    const seconds = date.getSeconds() < 10 ? "0"+ date.getSeconds() : date.getSeconds();
    clock.innerText = `${hours}:${minutes}:${seconds}`;
  }


//사용자 이름 입력, 저장, 출력
  // 사용자 입력받은 이름 저장
  function saveName(userName){
    localStorage.setItem(userStorage, userName);
  }

  // 사용자 이름을 입력 받고 저장
  function submitName(){
    event.preventDefault();
    const userName = input.value;
    printGreetings(userName);
    saveName(userName);
  }

  //폼을 보여 줌
  function printNameForm(){
    form.classList.add(visible);
    form.addEventListener("submit", submitName);
  }

  // 인사말과 이름을 보여줌
  function printGreetings(userName){
    form.classList.remove(visible);
    greetings.classList.add(visible);
    greetings.innerText = `Hi ${userName}`;
  }

  //이름을 받았는지 안 받았는지 체크
  function loadUser(){
    const currentUser = localStorage.getItem(userStorage);
    if(currentUser === null){
      printNameForm();
    }else{
      printGreetings(currentUser);
    }
  }

//할일 입력, 저장, 리스트 출력
  function printToDoList(toDoItem){
    const li = document.createElement('li'),
          delBtn = document.createElement('button'),
          toDoText = document.createElement('span'),
          toDoId = document.createElement('span');
    let toDoIdValue = 1 + toDoArray.length;
    toDoId.innerText = toDoIdValue;
    toDoText.innerText = toDoItem;
    delBtn.innerText = 'delete';
    li.appendChild(toDoId);
    li.appendChild(toDoText);
    li.appendChild(delBtn);
    toDoList.appendChild(li);
    const toDoObj = {
      text : toDoItem,
      id : toDoIdValue,
    }
    toDoArray.push(toDoObj);
    localStorage.setItem(toDoStorage, JSON.stringify(toDoArray));
  }

  function submitToDo(){
    event.preventDefault();
    const toDoItem = toDoInput.value;
    printToDoList(toDoItem);
    console.log(toDoItem);
    toDoInput.value = "";
  }

  function loadToDoList(){
    const currentToDoList = localStorage.getItem(toDoStorage);
    if(currentToDoList !== null){
      printToDoList();
      console.log(currentToDoList);
    }
    toDoForm.addEventListener("submit", submitToDo);
  }

  function init(){
    getTime();
    setInterval(getTime, 1000);
    loadUser();
    loadToDoList();
  }
  init();
}
