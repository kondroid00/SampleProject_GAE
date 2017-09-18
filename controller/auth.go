package controller

import (
	"dto"
	"encoding/json"
	"model"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func Login(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	usersModel := &model.UsersModel{
		Ctx: ctx,
	}
	log.Debugf(ctx, ("login = "))

	params := &struct {
		UserId string `json:"userId"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(params); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	user, err := usersModel.Login(params.UserId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}

	authModel := &model.AuthModel{
		Ctx: ctx,
	}
	token, err := authModel.CreateToken()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	result := &struct {
		User  dto.User  `json:"user"`
		Token dto.Token `json:"token"`
	}{
		User:  *user,
		Token: *token,
	}

	returnJson(&w, result)
}
