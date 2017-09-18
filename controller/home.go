package controller

import (
	"fmt"
	"model"
	"net/http"

	"google.golang.org/appengine"
)

func HomeIndex(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	roomsModel := model.NewRoomsModel(&ctx)

	rooms, err := roomsModel.FetchRooms()
	if err != nil {
		fmt.Fprint(w, "fetch error")
		return
	}

	returnJson(&w, rooms)
}
