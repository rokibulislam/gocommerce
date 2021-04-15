package domain

type Package struct {
	Id          int    `json: id xml: "id"`
	Name        string `json: name xml: "name"`
	Quantity    string `json: qty xml: "qty"`
	Price       string `json: price xml: "price"`
	Description string `json: description xml: "description"`
}
