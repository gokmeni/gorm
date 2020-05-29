package entity

type OrderDetail struct {
	OrderID   int     `gorm:"column:order_id"`
	ProductID int     `gorm:"column:product_id"`
	UnitPrice float64 `gorm:"column:unit_price"`
	Quantity  int     `gorm:"column:quantity"`
	Discount  float64 `gorm:"column:discount"`
}

func (OrderDetail) TableName() string {
	return "order_details"
}
