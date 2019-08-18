package controllers

import (
	"api/database"
	"api/models"
	"encoding/json"
	"fmt"
	"net/http"
)

type Product models.Product

func (p *Product) GetAllProduct(w http.ResponseWriter, r *http.Request) {
	var collection database.Collection
	products := collection.GetCollection(database.COLLECTION_PRODUCT)
	json.NewEncoder(w).Encode(products)
}
func (p *Product) InsertProduct(w http.ResponseWriter, r *http.Request) {
	var collection database.Collection
	products := collection.GetCollection(database.COLLECTION_PRODUCT)
	fmt.Println(r.Body)
	fmt.Println(json.NewDecoder(r.Body).Decode(p))
	fmt.Println(p)
	products.Insert(Product{
		Name:  p.Name,
		Price: p.Price,
	})
}
func (p *Product) GetProductById(w http.ResponseWriter, r *http.Request) {

}
func (p *Product) UpdateProduct(w http.ResponseWriter, r *http.Request) {

}
func (p *Product) DeleteProduct(w http.ResponseWriter, r *http.Request) {

}
