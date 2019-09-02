package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Index per passione
func Index(w http.ResponseWriter, r *http.Request) {
	fileBin, err := ioutil.ReadFile("index.html")

	if err != nil {
		fmt.Fprintf(w, "Impossibile leggere file")
	}

	htmlContent := string(fileBin)

	fmt.Fprintf(w, htmlContent)
}

func GetAllComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(comments)
}

//Data è un elemento per trasferire dati da web
type AddCommentRequest struct {
	Message     string
	AuthorEmail string
	AuthorName  string
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func AddComment(w http.ResponseWriter, r *http.Request) {
	var req AddCommentRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeErrorResponse(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if len(req.Message) < 3 || req.AuthorEmail == "" || len(req.AuthorName) < 3 {
		writeErrorResponse(w, "Invalid data", http.StatusBadRequest)
	}

	a := NewAuthor(req.AuthorEmail, req.AuthorName)
	c := NewComment(a, req.Message)

	comments = append(comments, c)
	fmt.Print("Aggiunto messaggio da" + req.AuthorEmail)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(c)
}

func writeErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	e := ErrorResponse{message}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(e)
}
