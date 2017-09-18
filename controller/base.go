package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

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
