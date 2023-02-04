package controllers

import (
	"net/http"

	"github.com/8tako8tako8/go_todo/config"
)

func StartMainServer() error {
    http.HandleFunc("/", top)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
