package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestPostAndGetHandlersValidPayload(t *testing.T) {
	/* POST a valid message to the postHandler(). */

	payload := []byte(`{"message": "foo"}`)
	req := httptest.NewRequest("POST", "/messages", bytes.NewReader(payload))
	w := httptest.NewRecorder()

	postHandler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	strBody := string(body)

	expectedHash := "2c26b46b68ffc68ff99b453c1d30413413422d706483bfa0f98a5e886266e7ae"
	expectedBody := fmt.Sprintf("{\n  \"digest\": \"%s\"\n}\n", expectedHash)

	if strBody != expectedBody {
		t.Errorf("Expected: '%s'\nGot: %s", expectedBody, strBody)
	}

	/* Now, lets see if we can GET our message back by its hash. */

	req = httptest.NewRequest("GET", "/messages/"+expectedHash, nil)
	w = httptest.NewRecorder()

	getHandler(w, req)

	resp = w.Result()
	body, _ = ioutil.ReadAll(resp.Body)
	strBody = string(body)

	expectedBody = "{\n  \"message\": \"foo\"\n}\n"

	if strBody != expectedBody {
		t.Errorf("Expected: %s,\nGot: %s", strBody, expectedBody)
	}
}

func TestPostHandlerJsonUnmarshalError(t *testing.T) {
	payload := []byte(`{"bad json`)
	req := httptest.NewRequest("POST", "/messages", bytes.NewReader(payload))
	w := httptest.NewRecorder()

	postHandler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	strBody := string(body)

	expectedStatusCode := 400

	if resp.StatusCode != expectedStatusCode {
		t.Errorf("Expected StatusCode: %d, Got: %d", expectedStatusCode, resp.StatusCode)
	}

	expectedBody := "{\n  \"error\": \"Unmarshal Error: 'unexpected end of JSON input' with Body: '{\\\"bad json'\"\n}\n"

	if strBody != expectedBody {
		t.Errorf("Expected: '%s'\nGot: %s", expectedBody, strBody)
	}
}

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
