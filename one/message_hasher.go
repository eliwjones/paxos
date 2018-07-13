package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type payload struct {
	Digest  string `json:"digest,omitempty"`
	Error   string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}

// TODO: make sure to never restart the service or we lose all of our data.
var db = map[string]string{}

func main() {
	http.HandleFunc("/messages/", getHandler)
	http.HandleFunc("/messages", postHandler)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal(fmt.Errorf("$PORT is not set"))
	}

	http.ListenAndServe(":"+port, nil)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.SetIndent("", "  ")

	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		jsonEncoder.Encode(payload{Error: "I only accept POST requests with application/json data of the form: {'message': 'foo'}."})
		return
	}

	b := bytes.NewBuffer(nil)
	_, err := io.Copy(b, r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		jsonEncoder.Encode(payload{Error: err.Error()})
		return
	}

	p := payload{}
	json.Unmarshal(b.Bytes(), &p)

	messageHash := fmt.Sprintf("%x", sha256.Sum256([]byte(p.Message)))

	db[messageHash] = p.Message

	jsonEncoder.Encode(payload{Digest: messageHash})
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.SetIndent("", "  ")

	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		jsonEncoder.Encode(payload{Error: "I only accept GET requests."})
		return
	}

	splitURL := strings.Split(r.URL.Path, "/messages/")
	messageHash := splitURL[1]

	if message, found := db[messageHash]; found {
		jsonEncoder.Encode(payload{Message: message})
	} else {
		w.WriteHeader(http.StatusNotFound)
		jsonEncoder.Encode(payload{Error: "Message not found."})
	}
}
