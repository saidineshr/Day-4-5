package Controllers

import (
	"awesomeProject3/Models"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type response struct {
	//p *Models.Transaction
	message string
}

func CreateTransaction(c *gin.Context) {
	//
	var r response
	var transaction Models.Transaction
	transaction.DateTime = time.Now()
	var prod Models.Product
	err := c.BindJSON(&transaction)

	//if err != nil {
	//	return
	//}
	prodId := strconv.Itoa(transaction.ProductId)

	var t Models.Transaction

	err = Models.GetBYTransactionProductId(&t, prodId)
	if err == nil {
		fmt.Println("kk",transaction.DateTime.Unix() - t.DateTime.Unix())
		if (transaction.DateTime.Unix() - t.DateTime.Unix()) < 300 {
			c.AbortWithStatusJSON(http.StatusMethodNotAllowed, gin.H{"message": "Cannot fetch as the quantity...Wit for few more minutes"})
			return
		}
	}


	err = Models.GetByProductID(&prod, prodId)

	if prod.Quantity < (&transaction).ProductQuantity {
		c.AbortWithStatusJSON(http.StatusMethodNotAllowed, gin.H{"message": "Cannot fetch as the quantity is less"})
		return
	}else{
		prod.Quantity-=(&transaction).ProductQuantity
		Models.UpdateProductQuantity(&prod,prodId)
	}

	err = Models.CreateTransaction(&transaction)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
		return
	} else {

		r.message = "Success"
		resp:= struct {
			Models.Transaction
			response
		}{transaction,r}
		fmt.Println(resp)
		c.JSON(http.StatusOK, resp)
	}
}

func GetTransactionByRetailerID(c *gin.Context) {
	id := c.Params.ByName("id")
	var transaction Models.Transaction
	err := Models.GetByTransactionRetailerID(&transaction, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, transaction)
	}
}
func GetTransactionByCustomerID(c *gin.Context) {
	id := c.Params.ByName("id")
	var transaction Models.Transaction
	err := Models.GetByTransactionCustomerID(&transaction, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, transaction)
	}
}
