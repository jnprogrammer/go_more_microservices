package data

import "time"

// Product defines the structure for an APU product
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json: "name"`
	Description string  `json: "description"`
	Price       float32 `json:"price"`
	SKU         string  `json: "sku"`
	CreatedOn   string  `json:"-"`
	UpdateOn    string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

func GetProcuts() []*Product {
	return productList
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdateOn:    time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee sans milk",
		Price:       1.60,
		SKU:         "gw42",
		CreatedOn:   time.Now().UTC().String(),
		UpdateOn:    time.Now().UTC().String(),
	},
}
