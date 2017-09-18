package main

import (
	"net/http"

	"controller"

	"github.com/gorilla/mux"
)

func init() {

	r := mux.NewRouter()

	// auth
	r.HandleFunc("/login", controller.Login)

	// user
	r.HandleFunc("/user", controller.UserIndex)
	r.HandleFunc("/user/create", controller.UserCreate)

	// home
	r.HandleFunc("/home", controller.HomeIndex)

	// room
	r.HandleFunc("/room", controller.RoomIndex)
	r.HandleFunc("/room/create", controller.RoomCreate)

	http.Handle("/", r)
}
