package domain

type User struct {
	Name string `json: full_name xml: "name"`
	City string `json: city xml:"city"`
}
