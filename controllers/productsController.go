package controllers

import (
	"encoding/json"
	"golang-shop/models"
	"golang-shop/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var CreateProduct = func(w http.ResponseWriter, r *http.Request) {
	product := &models.Product{}

	err := json.NewDecoder(r.Body).Decode(product)
	if err != nil {
		utils.Respond(w, utils.Message(false, "error while decode product"))
		return
	}
	product.CreatedAt = int(time.Now().Unix())
	resp := product.Create()
	utils.Respond(w, resp)
}

var GetProducts = func(w http.ResponseWriter, r *http.Request) {
	products := models.GetProducts()
	resp := utils.Message(true, "succsess")
	resp["data"] = products
	utils.Respond(w, resp)
}

var GetProduct = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.Respond(w, utils.Message(false, "There was error by id"))
		return
	}
	product := models.GetProduct(id)
	resp := utils.Message(true, "success")
	resp["data"] = product
	utils.Respond(w, resp)
}
