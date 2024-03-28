package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	if boby, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(boby), x); err != nil {
			return
		}
	}
}
