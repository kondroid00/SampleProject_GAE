package controller

import (
	"fmt"
	"net/http"
)

func TopIndex(w http.ResponseWriter, r *http.Request) {
	/*
		ctx := appengine.NewContext(r)
		m := &model.UsersModel{
			Ctx: ctx,
		}

		users, err := m.FetchUsers()
		if err != nil {

		}

		for i, _ := range users {
			log.Debugf(ctx, strconv.Itoa(i))
		}
	*/

	fmt.Fprint(w, "top")
}
