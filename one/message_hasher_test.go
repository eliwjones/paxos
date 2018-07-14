package main

import (
	"bytes"
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetHandlerNotFound(t *testing.T) {
	req := httptest.NewRequest("GET", "/messages/abcdefg", nil)
	w := httptest.NewRecorder()

	getHandler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	expectedStatusCode := 404

	if resp.StatusCode != expectedStatusCode {
		t.Errorf("Expected StatusCode: %d, Got: %d", expectedStatusCode, resp.StatusCode)
	}

	expectedBody := "{\n  \"error\": \"Message not found.\"\n}\n"

	if string(body) != expectedBody {
		t.Errorf("Expected: %s,\nGot: %s", string(body), expectedBody)
	}
}

func TestPostHandler(t *testing.T) {
	payload := []byte(`{"message": "foo"}`)
	req := httptest.NewRequest("POST", "/messages", bytes.NewReader(payload))
	w := httptest.NewRecorder()

	postHandler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	strBody := string(body)

	expectedHash := "2c26b46b68ffc68ff99b453c1d30413413422d706483bfa0f98a5e886266e7ae"

	if !strings.Contains(strBody, expectedHash) {
		t.Errorf("Expected: '%s' in body.\n\nGot: %s", expectedHash, strBody)
	}

	req = httptest.NewRequest("GET", "/messages/"+expectedHash, nil)
	w = httptest.NewRecorder()

	getHandler(w, req)

	resp = w.Result()
	body, _ = ioutil.ReadAll(resp.Body)

	expectedBody := "{\n  \"message\": \"foo\"\n}\n"

	if string(body) != expectedBody {
		t.Errorf("Expected: %s,\nGot: %s", string(body), expectedBody)
	}
}
