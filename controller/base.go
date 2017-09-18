package controller

import (
	"bytes"
	"dto"
	"encoding/json"
	"fmt"
	"model"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func checkToken(w http.ResponseWriter, r *http.Request) *dto.User {
	ctx := appengine.NewContext(r)

	params := &struct {
		Token string `json:"token"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(params); err != nil {
		log.Debugf(ctx, "decode error")
		log.Debugf(ctx, "%s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}

	log.Debugf(ctx, "userToken = %s", params.Token)

	authModel := model.NewAuthModel(&ctx)
	result, token := authModel.Verify(params.Token)
	switch result {
	case model.VERIFY_OK:
		usersModel := model.NewUsersModel(&ctx)
		user, err := usersModel.FetchUser(token.UserId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return user
	case model.VERIFY_NOT_FOUND:
		http.Error(w, "verify not found", http.StatusUnauthorized)
	case model.VERIFY_EXPIRED:
		http.Error(w, "verify expired", http.StatusUnauthorized)
	}

	return nil
}

func returnJson(w *http.ResponseWriter, v interface{}) {
	jsonBytes, err := json.Marshal(v)
	if err != nil {
		fmt.Fprint(*w, err)
		return
	}

	out := new(bytes.Buffer)
	json.Indent(out, jsonBytes, "", "    ")
	fmt.Fprint(*w, out.String())
}
