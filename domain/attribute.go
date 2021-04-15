package domain

type Attribute struct {
	Id   int    `json: id xml: "id"`
	Name string `json: name xml: "name"`
}
