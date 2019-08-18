package main

import (
	"api/controllers"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Response struct {
	Message    string `json:"message"`
	StatusCode string `json:"statusCode"`
}

func main() {
	route := mux.NewRouter()

	var productController controllers.Product
	http.Handle("/", route)
	route.HandleFunc("/", requestHandler)
	route.HandleFunc("/product", productController.GetAllProduct)
	route.HandleFunc("/insert", productController.InsertProduct)
	if err := http.ListenAndServe(":1111", nil); err != nil {
		panic(err)
	}
}
func requestHandler(w http.ResponseWriter, r *http.Request) {
	data := Response{StatusCode: "200", Message: "Server runinng"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
