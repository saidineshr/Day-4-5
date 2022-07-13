package Models

import "time"

type Product struct {
	Id         int    `json:"id" binding:"required"`
	Name       string `json:"name"binding:"required"`
	Price      int    `json:"price"binding:"required"`
	Quantity   int    `json:"quantity"`
	RetailerId int    `json:"retailer_id"binding:"required"`
}

type Transaction struct {
	CustomerId      int       `json:"customer_id"binding:"required"`
	TransactionId   string    `json:"transaction_id"binding:"required"`
	ProductId       int       `json:"product_id"binding:"required"`
	DateTime        time.Time `json:"date_time"binding:"required"`
	RetailerId      int       `json:"retailer_id"binding:"required"`
	ProductQuantity int       `json:"product_quantity"binding:"required"`
}

//type Retailer struct {
//	Id        int `json:"id"binding:"required"`
//	ProductId int `json:"product_id"binding:"required"`
//}

func (b *Product) TableName() string {
	return "product"
}

func (b *Transaction) TableName() string {
	return "transaction"
}

//func (b *Retailer) TableName() string {
//	return "retailer"
//}
