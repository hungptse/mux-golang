package controllers

import (
	"api/database"
	"api/models"
	"api/utils"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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
	resData := utils.Response(listPro, "OK!", "200", len(listPro))
	log.Println(resData)
	json.NewEncoder(w).Encode(resData)
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
	var collection database.Collection
	products := collection.GetCollection(database.COLLECTION_PRODUCT)
	vars := mux.Vars(r)
	// id := r.FormValue("id") // from query param
	id := vars["id"] // from path param
	var pro []Product
	err := products.Find(bson.M{"id": id}).All(&pro)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(pro)
}
func (p *Product) UpdateProduct(w http.ResponseWriter, r *http.Request) {

}
func (p *Product) DeleteProduct(w http.ResponseWriter, r *http.Request) {

}
