package handler

import (
	//	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostUser(t *testing.T) {
	req, err := http.NewRequest("POST", "/v1/user", nil)
	if err != nil {
		t.Fatalf("fail to run server")
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PostUser)
	handler.ServeHTTP(rr, req)

	//	server := httptest.NewServer(handler)

	//	var server = httptest.NewServer(http.h)
	//	defer server.Close()

	t.Log(rr.Code)
}
