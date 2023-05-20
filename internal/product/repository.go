package product

import (
	"github.com/ContrerasJoel/gestor_go/db"
)

func ReadOne(id string) (Product, error) {
	var err error
	var product Product
	productFound := db.DB.First(&product, id)
	err = productFound.Error
	if err != nil {
		return product, err
	}
	return product, nil
}

func ReadAll() ([]Product, int64) {
	var r int64
	var products []Product
	productsFound := db.DB.Find(&products)
	r = productsFound.RowsAffected
	if r == 0 {
		return nil, r
	}
	return products, r
}

func CreateOne(product *Product) error {
	var err error
	saveProduct := db.DB.Create(&product)
	err = saveProduct.Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateOne(product Product, id string) (Product, error) {
	productFound, err := ReadOne(id)
	if err != nil {
		return productFound, err
	}
	productFound.Name = product.Name
	productFound.Sku = product.Sku
	productFound.Price = product.Price
	productFound.Color = product.Color
	productFound.Measures = product.Measures
	productFound.Material = product.Material
	productFound.Category = product.Category
	productFound.Quantity = product.Quantity
	productFound.Description = product.Description
	productFound.Images = product.Images
	db.DB.Save(&productFound)
	return productFound, nil
}

func DeleteOne(id string) int64 {
	var r int64
	var product Product
	deleteProduct := db.DB.Unscoped().Delete(&product, id)
	r = deleteProduct.RowsAffected
	if r == 0 {
		return r
	}
	return r
}
