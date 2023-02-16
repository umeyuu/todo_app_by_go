package main

import (
	"fmt"
	"sqlite/go/src/todo_app_by_go/app/controllers"
	"sqlite/go/src/todo_app_by_go/app/models"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainSever()
}
