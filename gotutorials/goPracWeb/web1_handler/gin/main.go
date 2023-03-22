package main

import (
	"./myapp"
)

func main() {
	router := myapp.NewHttpHandler()
	router.Run("1234")
}
