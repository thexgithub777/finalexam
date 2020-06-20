package customer

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello",
	})
}

//Router root router
func Router() *gin.Engine {
	CreateTable()
	r := gin.Default()

	apiGrp := r.Group("/")

	apiGrp.Use(AuthorizationToken)

	apiGrp.GET("customers", GetAllCustomersHandler)
	apiGrp.GET("customers/:id", GetCustomerIDHandler)
	apiGrp.POST("customers", CreateCustomerHandler)
	apiGrp.PUT("customers/:id", UpdateCustomerHandler)
	apiGrp.DELETE("customers/:id", DeleteCustomerHandler)

	return r
}
