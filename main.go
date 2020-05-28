package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gokmeni/gorm/ds"
	"github.com/gokmeni/gorm/entity"
	"github.com/jinzhu/gorm"
	"net/http"
)

const (
	httpServerPort string = ":6789"
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

	defer closeConnection(db)

	closeConnection(db)

	startServer()
}

func closeConnection(db *gorm.DB) {
	err := db.Close()

	if err != nil {
		fmt.Printf("connection close error: %v\n", err)
	}
}

func startServer() {
	router := gin.Default()

	router.GET(getCustomerByIdRoute, getCustomerByID)

	err := router.Run(httpServerPort)

	if err != nil {
		fmt.Printf("server could not be started -> %v\n", err)
	}
}

func getCustomerByID(c *gin.Context) {
	customerID := c.Param(getCustomerByIdRouteParameter)

	var customer entity.Customer

	err := db.Where(&entity.Customer{CustomerID: customerID}).First(&customer).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.String(http.StatusNotFound, getCustomerByIdNotFoundError)
		} else {
			fmt.Printf("connection close error -> %v\n", err)
			c.String(http.StatusForbidden, getCustomerByIdSystemError)
		}
	} else {
		c.JSON(http.StatusOK, customer)
	}
}
