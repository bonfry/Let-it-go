package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var comments []*Comment

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", Index).Methods("GET")
	r.HandleFunc("/comments", GetAllComments).Methods("GET")
	r.HandleFunc("/comments", AddComment).Methods("POST")

	autore := NewAuthor("bonfry@bonfry.com", "Bonfry")
	c1 := NewComment(autore, "BellaCiao")
	comments = append(comments, c1)

	c2 := NewComment(autore, "Portami a Los Angeles")
	comments = append(comments, c2)

	fmt.Println(comments)
	fmt.Println("Ho acceso il server")

	panic(http.ListenAndServe(":8800", r))
}
