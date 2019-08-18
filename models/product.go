package models

//Product model
type Product struct {
	ID    string `bson:"id" json:"id"`
	Name  string `bson:"name" json:"name"`
	Price int    `bson:"price" json:"price"`
}
