package Model

import "time"

type Discount struct {
	Id              int       `json:"id" gorm:"type INT(10) UNSIGNED NOT NULL AUTO_INCREMENT;primary key"`
	Qty             int       `json:"qty"`
	Type            string    `json:"type"`
	Result          int       `json:"result"`
	ExpiredAt       int       `json:"expired_at"`
	ExpiredAtFormat string    `json:"expired_at_format"`
	StringFormat    string    `json:"string_format"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"Updated_at"`
}
