package controllers

import (
	"net/http"

	"github.com/wd30gsrc/go-simple-todo-app/config"
)

func StartMainServer() error {
	http.HandleFunc("/", top)
	return http.ListenAndServe(":" + config.Config.Port, nil)
}
