package main

type Author struct {
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}

func NewAuthor(email, fullname string) *Author {
	// devono essere messi in ordine di def in struct altrimenti [chiave valore],...
	return &Author{email, fullname}
}
