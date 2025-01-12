package Model

import "time"

type Cashier struct {
	Id        uint      `json:"id"`
	Name      string    `json:"id"`
	Passcode  string    `json:"passcode"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"Updated_at"`
}
