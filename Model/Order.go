package Model

import "time"

type Order struct {
	Id            int       `json:"id" gorm:"type INT(10) UNSIGNED NOT NULL AUTO_INCREMENT;primary key"`
	CashierId     int       `json:"cashier_id"`
	PaymentTypeId int       `json:"payment_type_id"`
	TotalPrice    int       `json:"total_price"`
	TotalPaid     int       `json:"total_paid"`
	TotalReturn   int       `json:"total_return"`
	ReceiptId     string    `json:"receipt_id"`
	IsDownload    string    `json:"is_download"`
	ProductId     string    `json:"product_id"`
	Quantities    string    `json:"quantities"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"Updated_at"`
}

type ProductResponseOrder struct {
	ProductId        int      `json:"productId" gorm:"type:INT(10) UNSIGNED NOT NULL AUTO_INCREMENT;primaryKey"`
	Name             string   `json:"name"`
	Price            int      `json:"price"`
	Qty              int      `json:"qty"`
	Discount         Discount `json:"discount"`
	TotalNormalPrice int      `json:"totalNormalPrice"`
	TotalFinalPrice  int      `json:"totalFinalPrice"`
}

type ProductOrder struct {
	Id         int    `json:"Id" gorm:"type:INT(10) UNSIGNED NOT NULL AUTO_INCREMENT;primaryKey"`
	Sku        string `json:"sku"`
	Name       string `json:"name"`
	Stock      int    `json:"stock"`
	Price      int    `json:"price"`
	Image      string `json:"image"`
	CategoryId int    `json:"categoryId"`
	DiscountId int    `json:"discountId"`
}

type RevenueResponse struct {
	PaymentTypeId int    `json:"paymentTypeId"`
	Name          string `json:"name"`
	Logo          string `json:"logo"`
	TotalAmount   int    `json:"totalAmount"`
}

type SoldResponse struct {
	ProductId   int    `json:"productId"`
	Name        string `json:"name"`
	TotalQty    int    `json:"totalQty"`
	TotalAmount int    `json:"totalAmount"`
}
