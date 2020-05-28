package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gokmeni/gorm/ds"
	"github.com/gokmeni/gorm/entity"
	"github.com/jinzhu/gorm"
	"net/http"
)

const (
	httpServerPort string = ":8080"
)

const (
	getCustomerByIdRoute          string = "/api/customer/:customerID"
	getCustomerByIdRouteParameter string = "customerID"
	getCustomerByIdNotFoundError  string = "Customer Not Found."
	getCustomerByIdSystemError    string = "System Error."
)

var db *gorm.DB

func main() {

	db = ds.GetConnection()

	defer db.Close()

	startServer()
}

func startServer() {
	router := gin.Default()

	router.GET(getCustomerByIdRoute, getCustomerByID)

	router.Run(httpServerPort)
}

func getCustomerByID(c *gin.Context) {
	customerID := c.Param(getCustomerByIdRouteParameter)

	var customer entity.Customer

	err := db.Debug().Where(&entity.Customer{CustomerID: customerID}).First(&customer).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.String(http.StatusNotFound, getCustomerByIdNotFoundError)
		} else {
			c.String(http.StatusForbidden, getCustomerByIdSystemError)
		}
	} else {
		c.JSON(http.StatusOK, customer)
	}
}
