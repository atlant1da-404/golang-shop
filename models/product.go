package models

import (
	"golang-shop/utils"
	"time"
)

type Product struct {
	Id          int    `sql:"id" json:"id"`
	Title       string `sql:"title" json:"title"`
	Description string `sql:"description" json:"description"`
	CreatedAt   int    `sql:"created_at"`
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

func (product *Product) SetTimeStamp() {
	product.CreatedAt = int(time.Now().Unix())
}
