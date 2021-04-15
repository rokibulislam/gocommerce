package domain

type Tag struct {
	Id          int    `json: id xml: "id"`
	Name        string `json: name xml: "name"`
	Slug        string `json: slug xml:"slug"`
	Description string `json: description xml:"description"`
}
