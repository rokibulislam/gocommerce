package domain

type Order struct {
	Id   int    `json: id xml: "id"`
	Name string `json: name xml: "name"`
	City string `json: city xml:"city"`
}
