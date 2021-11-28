package models

import (
	"fmt"
	"golang-shop/utils"
)

type Product struct {
	tableName   struct{} `sql:"product" pg:"product" json:"-"`
	Id          int      `sql:"id" json:"id"`
	Title       string   `sql:"title" json:"title"`
	Description string   `sql:"description" json:"description"`
	CreatedAt   int      `sql:"created_at"`
}

func (product *Product) Validate() (map[string]interface{}, bool) {

	if product.Title == "" {
		return utils.Message(false, "product is required"), true
	}

	if product.Description == "" {
		return utils.Message(false, "description is required"), true
	}

	return utils.Message(true, "success"), true
}

func (product *Product) Create() map[string]interface{} {
	if resp, ok := product.Validate(); !ok {
		return resp
	}

	GetDB().Model(product).Insert()
	resp := utils.Message(true, "success")
	resp["product"] = product
	return resp
}

func GetProducts() interface{} {
	var products []Product
	err := GetDB().Model(&products).Select()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return products
}
