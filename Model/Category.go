package Model

import "time"

type Category struct {
	Id        int       `json:"id" gorm:"type INT(10) UNSIGNED NOT NULL AUTO_INCREMENT;primary key"`
	Name      int       `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"Updated_at"`
}
