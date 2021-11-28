package main

import (
	"fmt"
	"log"

	"github.com/wd30gsrc/go-simple-todo-app/config"
)

func main() {
	fmt.Println(config.Config.Port)
	fmt.Println(config.Config.SQLDriver)
	fmt.Println(config.Config.DbName)
	fmt.Println(config.Config.LogFile)

	log.Println("test2")
}
