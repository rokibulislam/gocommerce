package domain

type Payment struct {
	Name string `json: full_name xml: "name"`
	City string `json: city xml:"city"`
}
