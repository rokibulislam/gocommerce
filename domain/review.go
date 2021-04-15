package domain

type Review struct {
	Id      int    `json: id xml: "id"`
	Comment string `json: comment xml: "comment"`
	Rating  string `json: rating xml:"rating"`
}
