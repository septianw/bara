package handler

import "net/http"

func NotFound(w http.ResponseWriter, r *http.Request) {
	//	w.WriteHeader(404)
	w.Header().Set("ContentType", "application/json")
	Send(w, OK, JSON, `{ "status": "I'am alive" }`)
}
