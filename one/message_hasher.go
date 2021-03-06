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
	"sync"
)

type payload struct {
	Digest  string `json:"digest,omitempty"`
	Error   string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}

// TODO: make sure to never restart the service or we lose all of our data.
// TODO: use Redis or some lightning fast key-value store for persistence.
var db = sync.Map{}

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
		errorMessage := "I only accept POST requests with application/json data of the form: {'message': 'foo'}."
		jsonEncoder.Encode(payload{Error: errorMessage})

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
	err = json.Unmarshal(b.Bytes(), &p)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorMessage := fmt.Sprintf("Unmarshal Error: '%s' with Body: '%s'", err.Error(), b.String())
		jsonEncoder.Encode(payload{Error: errorMessage})

		return
	}

	messageHash := fmt.Sprintf("%x", sha256.Sum256([]byte(p.Message)))

	db.Store(messageHash, p.Message)

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

	if message, found := db.Load(messageHash); found {
		jsonEncoder.Encode(payload{Message: message.(string)})
	} else {
		w.WriteHeader(http.StatusNotFound)
		jsonEncoder.Encode(payload{Error: "Message not found."})
	}
}
