package model

import (
	"dto"

	"github.com/mjibson/goon"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

// Model
type UsersModel struct {
	Ctx  context.Context
	kind string
}

func NewUsersModel(ctx *context.Context) *UsersModel {
	return &UsersModel{
		Ctx:  *ctx,
		kind: "User",
	}
}

func (u *UsersModel) Login(userId string) (*dto.User, error) {
	g := goon.FromContext(u.Ctx)

	user := &dto.User{
		Id: userId,
	}

	if err := g.Get(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UsersModel) FetchUsers() ([]dto.User, error) {
	g := goon.FromContext(u.Ctx)

	q := datastore.NewQuery(u.kind).Limit(10)

	users := make([]dto.User, 0, 10)
	if _, err := g.GetAll(q, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UsersModel) CreateUser(user *dto.User) (*datastore.Key, error) {
	g := goon.FromContext(u.Ctx)
	uId := getUUID()
	log.Debugf(u.Ctx, ("user uuid = " + uId.String()))
	user.Id = uId.String()
	return g.Put(user)
}
