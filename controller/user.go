package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"dto"
	"model"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func UserIndex(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	usersModel := model.NewUsersModel(&ctx)

	users, err := usersModel.FetchUsers()
	if err != nil {
		fmt.Fprint(w, "fetch error")
		return
	}

	log.Debugf(ctx, "fetch")
	returnJson(&w, users)
}

func UserCreate(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	usersModel := model.NewUsersModel(&ctx)

	params := &struct {
		Name string `json:"name"`
	}{}

	//log.Debugf(ctx, "body = %s", helper.GetBodyParamString(r))

	if err := json.NewDecoder(r.Body).Decode(params); err != nil {
		log.Debugf(ctx, "decode error")
		log.Debugf(ctx, "%s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user := &dto.User{
		Name: params.Name,
	}
	_, err := usersModel.CreateUser(user)
	if err != nil {
		fmt.Fprint(w, "create error")
		log.Debugf(ctx, "create error")
		return
	}

	authModel := model.NewAuthModel(&ctx)
	token, err := authModel.CreateToken(user.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
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

func UserShow(w http.ResponseWriter, r *http.Request) {

}

/*
func UserEdit(w http.ResponseWriter, r *http.Request) {

}
*/

func UserUpdate(w http.ResponseWriter, r *http.Request) {

}

func UserDelete(w http.ResponseWriter, r *http.Request) {

}
