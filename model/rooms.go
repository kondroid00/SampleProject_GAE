package model

import (
	"dto"

	"golang.org/x/net/context"

	"github.com/mjibson/goon"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

// Model
type RoomsModel struct {
	Ctx  context.Context
	kind string
}

func NewRoomsModel(ctx *context.Context) *RoomsModel {
	return &RoomsModel{
		Ctx:  *ctx,
		kind: "Room",
	}
}

func (u *RoomsModel) FetchRooms() ([]dto.Room, error) {
	g := goon.FromContext(u.Ctx)

	q := datastore.NewQuery(u.kind).Limit(10)

	rooms := make([]dto.Room, 0, 10)
	if _, err := g.GetAll(q, &rooms); err != nil {
		return nil, err
	}

	return rooms, nil
}

func (u *RoomsModel) CreateRoom(room *dto.Room) (*datastore.Key, error) {
	g := goon.FromContext(u.Ctx)
	uId := getUUID()
	log.Debugf(u.Ctx, ("room uuid = " + uId.String()))
	room.Id = uId.String()
	return g.Put(room)
}
