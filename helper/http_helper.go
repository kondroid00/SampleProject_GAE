package helper

import (
	"io/ioutil"
	"net/http"
)

func GetBodyParamString(r *http.Request) string {
	b, err1 := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err1 != nil {

	}
	return string(b)
}
