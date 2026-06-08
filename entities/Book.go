package entities

type Book struct {
	Id              int    `json:"id"`
	Title           string `json:"title"`
	AuthorId        int    `json:"authorId"`
	Genre           string `json:"genre"`
	PublicationDate string `json:"publicationDate"`
}
