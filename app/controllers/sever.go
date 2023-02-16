package controllers

import (
	"net/http"
	"sqlite/go/src/todo_app_by_go/config"
)

func StartMainSever() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	http.HandleFunc("/", top)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
