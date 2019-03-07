package handler

import "net/http"

func Version(w http.ResponseWriter, r *http.Request) {
	Send(w, OK, JSON, `{ "version": "0.1.0" }`)
}
