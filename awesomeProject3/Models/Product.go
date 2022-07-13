package Models

import (
	"awesomeProject3/Config"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// GetAllProducts Fetch all product data
func GetAllProducts(product *[]Product) (err error) {
	if err = Config.DB.Find(product).Error; err != nil {
		return err
	}
	return nil
}

// CreateProduct ... Insert New data
func CreateProduct(product *Product) (err error) {
	if err = Config.DB.Create(product).Error; err != nil {
		return err
	}
	return nil
}

// GetProductByID ... Fetch only one Product by Id
func GetByProductID(product *Product, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(product).Error; err != nil {
		return err
	}
	return nil
}

// UpdateProduct ... Update Product
func UpdateProductPrice(product *Product, id string) (err error) {
	fmt.Println(product)
	Config.DB.Save(product)
	return nil
}
func UpdateProductQuantity(product *Product, id string) (err error) {
	fmt.Println(product)
	Config.DB.Save(product)
	return nil
}

// DeleteProduct ... Delete Product
func DeleteProduct(product *Product, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(product)
	return nil
}
