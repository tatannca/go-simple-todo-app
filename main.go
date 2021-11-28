package main

import (
	"fmt"

	"github.com/wd30gsrc/go-simple-todo-app/app/controllers"
	"github.com/wd30gsrc/go-simple-todo-app/app/models"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()
}
