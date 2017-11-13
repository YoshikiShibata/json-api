package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestListTasksHandler(t *testing.T) {
	router := httprouter.New()
	router.GET("/", ListTasks)

	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	s := Server{router}
	s.ServeHTTP(w, r)

	expectedContentTypeHeader := "application/json; charset=utf-8"
	if w.Header().Get("Content-Type") != expectedContentTypeHeader {
		t.Errorf("handler returned unexpected Content-Type header: got %v, but want %v", w.Header().Get("Content-Type"), expectedContentTypeHeader)
	}

	expectedCode := http.StatusOK
	if w.Code != expectedCode {
		t.Errorf("handler returned unexpected status code: got %v, but want %v", w.Code, expectedCode)
	}
}
