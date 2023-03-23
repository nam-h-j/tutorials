package main

import (
	"./myapp"
)

func main() {
	myapp.NewHandler().Run(":1234")
}
