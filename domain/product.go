package domain

type Product struct {
	Id          int    `json: id xml: "id"`
	Title       string `json: title xml: "title"`
	Description string `json: description xml:"description"`
	Price       string `json: price xml:"price"`
}
