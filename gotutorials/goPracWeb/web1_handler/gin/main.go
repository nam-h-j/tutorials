package main

import (
	"./myapp"
)

func main() {
	myapp.NewHttpHandler().Run("1234")
}
