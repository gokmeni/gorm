package entity

type Customer struct {
	CustomerID   string `gorm:"primary_key;column:customer_id"`
	CompanyName  string `gorm:"column:company_name"`
	ContactName  string `gorm:"column:contact_name"`
	ContactTitle string `gorm:"column:contact_title"`
	Address      string `gorm:"column:address"`
	City         string `gorm:"column:city"`
	Region       string `gorm:"column:region"`
	PostalCode   string `gorm:"column:postal_code"`
	Country      string `gorm:"column:country"`
	Phone        string `gorm:"column:phone"`
	Fax          string `gorm:"column:fax"`

	Orders []Order `gorm:"foreignkey:CustomerID"`
}

func (Customer) TableName() string {
	return "customers"
}
