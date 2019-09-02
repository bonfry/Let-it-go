package main

// Comment è la struttura dei commenti
type Comment struct {
	Message string  `json:"message"` // esempio di tag struttura
	Author  *Author `json:"author"`  // puoi accederci con reflection(?)
}

//NewComment crea un nuovo commento
func NewComment(a *Author, m string) *Comment {
	return &Comment{m, a}
}
