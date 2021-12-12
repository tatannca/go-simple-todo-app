package main

import (
	"fmt"

	"github.com/wd30gsrc/go-simple-todo-app/app/controllers"
	"github.com/wd30gsrc/go-simple-todo-app/app/models"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()
	

	// user, _ := models.GetUserByEmail("test@example.com")
	// fmt.Println(user)

	// session, err := user.CreateSession()
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(session)

	// valied, _ := session.CheckSession()
	// fmt.Println(valied)
}
