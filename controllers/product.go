package controllers

import (
	"api/database"
	"api/models"
	"encoding/json"
	"log"
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

type Product models.Product

func (p *Product) GetAllProduct(w http.ResponseWriter, r *http.Request) {
	var collection database.Collection
	products := collection.GetCollection(database.COLLECTION_PRODUCT)
	var listPro []Product
	err := products.Find(bson.M{}).All(&listPro)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(listPro)
}
func (p *Product) InsertProduct(w http.ResponseWriter, r *http.Request) {
	var collection database.Collection
	products := collection.GetCollection(database.COLLECTION_PRODUCT)
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
