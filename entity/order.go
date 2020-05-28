package entity

import "time"

type Order struct {
	OrderID        int       `gorm:"primary_key;column:order_id"`
	CustomerID     string    `gorm:"column:customer_id"`
	EmployeeID     int       `gorm:"column:employee_id"`
	OrderDate      time.Time `gorm:"column:order_date"`
	RequiredDate   time.Time `gorm:"column:required_date"`
	ShippedDate    time.Time `gorm:"column:shipped_date"`
	ShipVia        int       `gorm:"column:ship_via"`
	Freight        float64   `gorm:"column:freight"`
	ShipName       string    `gorm:"column:ship_name"`
	ShipAddress    string    `gorm:"column:ship_address"`
	ShipCity       string    `gorm:"column:ship_city"`
	ShipRegion     string    `gorm:"column:ship_region"`
	ShipPostalCode string    `gorm:"column:ship_postal_code"`
	ShipCountry    string    `gorm:"column:ship_country"`
}

func (Order) TableName() string {
	return "orders"
}
