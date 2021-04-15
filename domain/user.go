package domain

type User struct {
	Id          string `json: id xml: "id"`
	Name        string `json: name xml: "name"`
	City        string `json: city xml:"city"`
	Zipcode     string `json: zipcode xml:"zipcode"`
	DateofBirth string `json: dateofbirth xml:"dateofbirth"`
	Email       string `json: email xml: "email"`
	Password    string `json: password xml: "password"`
	Status      string `json: status xml: "status"`
}

type JWT struct {
	Token string `json:"token"`
}

type Error struct {
	Message string `json:"message"`
}

type UserRepository interface {
	FinaAll() ([]User, error)
}
