package Model

import "time"

type Product struct {
	Id               int       `json:"id" gorm:"type INT(10) UNSIGNED NOT NULL AUTO_INCREMENT;primary key"`
	Sku              string    `json:"sku"`
	Name             string    `json:"name"`
	Stock            int       `json:"stock"`
	Price            int       `json:"price"`
	Image            string    `json:"receipt_id"`
	TotalFinalPrice  int       `json:"total_final_price"`
	TotalNormalPrice int       `json:"total_normal_price"`
	CategoryId       int       `json:"category_id"`
	DiscountId       int       `json:"discount_id"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"Updated_at"`
}
