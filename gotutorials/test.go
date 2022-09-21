package main

import "fmt"

func main() {
	//vars 변수
	var hello string = "world"
	var world = "hello"
	안뇽 := "반갑읍니다"
	fmt.Println(hello)
	fmt.Println(world)
	fmt.Println(안뇽)

	//if
	if 안뇽 == "반갑읍니다" {
		fmt.Println("저두요 ㅎㅎㅎ;")
	}

	//Arr 배열
	배열이 := [3]string{"저는", "배열", "입니다."}
	fmt.Println(배열이)

	//Dynamic Arr 동적배열
	dynamicArr := []string{"dynamic", "Arr", "I'm"}
	fmt.Println(dynamicArr)

	//Map 맵
	mapmap := make(map[string]string)
	mapmap["name"] = "김정국"
	fmt.Println(mapmap)

	//포인터
	pointer := "THIS IS POINTER!!!!"
	getPointer := &pointer
	fmt.Println(getPointer)
	fmt.Println(*getPointer)
}
