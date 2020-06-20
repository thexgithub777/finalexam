package customer

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//CreateCustomerHandler handler to create customer
func CreateCustomerHandler(c *gin.Context) {
	cust := Customer{}
	if err := c.ShouldBindJSON(&cust); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var err error
	cust.ID, err = CreateCustomer(&cust)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, cust)
}

//GetCustomerIDHandler handler to get customer by id
func GetCustomerIDHandler(c *gin.Context) {
	rqid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cust, err := FindCustomerByID(rqid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, cust)
}

//GetAllCustomersHandler handler to get all customers
func GetAllCustomersHandler(c *gin.Context) {
	custs, err := FindAllCustomers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, custs)
}

//UpdateCustomerHandler update customer by id
func UpdateCustomerHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	cust := &Customer{}
	if err := c.ShouldBindJSON(cust); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if id != cust.ID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id url and id in request not equal."})
		return
	}
	cust.ID = id

	err = UpdateCustomerByID(cust)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, cust)
}

//DeleteCustomerHandler delete customer by id
func DeleteCustomerHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err = DeleteCustomerByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "customer deleted"})
}
