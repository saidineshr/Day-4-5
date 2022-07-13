package Controllers

import (
	"awesomeProject3/Models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetProducts ... Get all Products
func GetProducts(c *gin.Context) {
	var Product []Models.Product
	err := Models.GetAllProducts(&Product)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, Product)
	}
}

// CreateProduct ... Create Product
func CreateProduct(c *gin.Context) {
	var Product Models.Product
	c.BindJSON(&Product)
	err := Models.CreateProduct(&Product)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, Product)
	}
}

// GetProductByID ... Get the Product by id
func GetProductByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var Product Models.Product
	err := Models.GetByProductID(&Product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, Product)
	}
}

// UpdateProduct ... Update the Product information
func UpdateProductPrice(c *gin.Context) {
	var Product Models.Product
	id := c.Params.ByName("id")
	err := Models.GetByProductID(&Product, id)
	if err != nil {
		c.JSON(http.StatusNotFound, Product)
	}
	c.BindJSON(&Product)
	err = Models.UpdateProductPrice(&Product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, Product)
	}
}

func UpdateProductQuantity(c *gin.Context) {
	var Product Models.Product
	id := c.Params.ByName("id")
	err := Models.GetByProductID(&Product, id)
	if err != nil {
		c.JSON(http.StatusNotFound, Product)
	}
	c.BindJSON(&Product)
	err = Models.UpdateProductQuantity(&Product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, Product)
	}
}

// DeleteProduct ... Delete the Product
func DeleteProduct(c *gin.Context) {
	var Product Models.Product
	id := c.Params.ByName("id")
	err := Models.DeleteProduct(&Product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
