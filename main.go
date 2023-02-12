package main

import (
	"fmt"
	"log"
	"sqlite/go/src/todo_app_by_go/config"
)

func main() {
	fmt.Println(config.Config.Port)
	fmt.Println(config.Config.SQLDriver)
	fmt.Println(config.Config.DbNmae)
	fmt.Println(config.Config.LogFile)

	log.Println("test")
}
