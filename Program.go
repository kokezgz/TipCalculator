package main

import (
	"./Controllers"
)

func main() {
	var controller Controllers.Controller
	controller.StartServer()
}
