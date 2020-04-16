package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	handler := setupHandler()
	recorder := httptest.NewRecorder()
	body := `{"repository":{"full_name":"c4s4/sweetohm"}}`
	request, _ := http.NewRequest("POST", "/", strings.NewReader(body))
	request.Header.Set("X-GitHub-Event", "push")
	handler.ServeHTTP(recorder, request)
	if recorder.Code != 200 {
		t.Errorf("%d != 200, %s\n", recorder.Code, recorder.Body)
	}
}
