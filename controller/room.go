package controller

import (
	"dto"
	"encoding/json"
	"fmt"
	"model"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func RoomIndex(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	currentUser := checkToken(w, r)
	if currentUser != nil {
		log.Debugf(ctx, "currentUser = %s", currentUser.Id)
	} else {
		log.Debugf(ctx, "currentUser = nil")
	}

	roomsModel := model.NewRoomsModel(&ctx)

	rooms, err := roomsModel.FetchRooms()
	if err != nil {
		fmt.Fprint(w, "fetch error")
		return
	}

	result := &struct {
		Rooms []dto.Room `json:"rooms"`
	}{
		Rooms: rooms,
	}

	returnJson(&w, result)
}

func RoomCreate(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	roomsModel := model.NewRoomsModel(&ctx)

	params := &struct {
		Name  string `json:"name"`
		Theme string `json:"theme"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(params); err != nil {
		log.Debugf(ctx, "decode error")
		log.Debugf(ctx, "%s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	room := &dto.Room{
		Name:  params.Name,
		Theme: params.Theme,
	}
	_, err := roomsModel.CreateRoom(room)
	if err != nil {
		fmt.Fprint(w, "create error")
		log.Debugf(ctx, "create error")
		return
	}

	result := &struct {
		Room dto.Room `json:"room"`
	}{
		Room: *room,
	}

	returnJson(&w, result)
}
