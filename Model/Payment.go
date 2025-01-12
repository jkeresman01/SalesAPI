package Model

import "time"

type Payment struct {
	Id            int       `json:"id" gorm:"type INT(10) UNSIGNED NOT NULL AUTO_INCREMENT;primary key"`
	Name          string    `json:"name"`
	Type          string    `json:"type"`
	PaymentTypeId int       `json:"payment_type_id"`
	Logo          string    `json:"logo"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"Updated_at"`
}
