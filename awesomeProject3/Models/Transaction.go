package Models

import (
	"awesomeProject3/Config"

	_ "github.com/go-sql-driver/mysql"
)

//GetAllProducts Fetch all product data
func GetAllTransactions(transaction *[]Transaction) (err error) {
	if err = Config.DB.Find(transaction).Error; err != nil {
		return err
	}
	return nil
}

//CreateProduct ... Insert New data
func CreateTransaction(transaction *Transaction) (err error) {
	if err = Config.DB.Create(transaction).Error; err != nil {
		return err
	}
	return nil
}

//GetProductByID ... Fetch only one Product by Id
func GetByTransactionRetailerID(transaction *Transaction, retailerId string) (err error) {
	if err = Config.DB.Where("retailer_id = ?", retailerId).First(transaction).Error; err != nil {
		return err
	}
	return nil
}

func GetByTransactionCustomerID(transaction *Transaction, customerId string) (err error) {
	if err = Config.DB.Where("customer_id = ?", customerId).First(transaction).Error; err != nil {
		return err
	}
	return nil
}

func GetBYTransactionProductId(transaction *Transaction,productId string)(err error){
	if err = Config.DB.Where("customer_id = ?", productId).First(transaction).Error; err != nil {
		return err
	}
	return nil
}
