package models

type OrderProduct struct {
	OrderID   int `json:"order_id" gorm:"primaryKey;autoIncrement:false"`
	ProductID int `json:"product_id" gorm:"primaryKey;autoIncrement:false"`
	Price     int `json:"price"`
	Quantity  int `json:"quantity"`
}
