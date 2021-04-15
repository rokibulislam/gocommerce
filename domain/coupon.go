package domain

type Coupon struct {
	Id     int    `json: id xml: "id"`
	Code   string `json: code xml: "code"`
	Amount string `json: amount xml:"amount"`
}
