package main

import (
	"fmt"

	"./cipher"
	"./lzw"
)

type Component interface {
	Operator(string)
}

var sentData string
var recvData string

// 데이터를 보내는 컴포넌트
type SendComponent struct{}

func (self *SendComponent) Operator(data string) {
	// Send data
	sentData = data
}

// 압축
type ZipComponent struct {
	com Component
}

func (self *ZipComponent) Operator(data string) {
	zipData, err := lzw.Write([]byte(data))
	if err != nil {
		panic(err)
	}
	self.com.Operator(string(zipData))
}

// 암호화
type EncryptComponent struct {
	key string
	com Component
}

func (self *EncryptComponent) Operator(data string) {
	encryptData, err := cipher.Encrypt([]byte(data), self.key)
	if err != nil {
		panic(err)
	}
	self.com.Operator(string(encryptData))
}

// 복호화
type DecryptComponent struct {
	key string
	com Component
}

func (self *DecryptComponent) Operator(data string) {
	decryptData, err := cipher.Decrypt([]byte(data), self.key)
	if err != nil {
		panic(err)
	}
	self.com.Operator(string(decryptData))
}

// 압축풀기
type UnzipComponent struct {
	com Component
}

func (self *UnzipComponent) Operator(data string) {
	unzipData, err := lzw.Read([]byte(data))
	if err != nil {
		panic(err)
	}
	self.com.Operator(string(unzipData))
}

type ReadComponent struct{}

func (self *ReadComponent) Operator(data string) {
	recvData = data
}

func main() {
	sender := &EncryptComponent{key: "force",
		com: &ZipComponent{
			com: &SendComponent{},
		},
	}

	sender.Operator("killing in the name of ?")
	fmt.Println(sentData)

	receiver := &UnzipComponent{
		com: &DecryptComponent{key: "force",
			com: &SendComponent{},
		},
	}

	receiver.Operator(sentData)
	fmt.Println(sentData)
}
