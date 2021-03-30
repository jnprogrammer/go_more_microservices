package data

import "testing"

func TestCheckValidation(t *testing.T) {
	p := &Product{
		Name:  "Tim",
		Price: 2.24,
		SKU:   "abs",
	}
	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}

//simple unit testing for validating Json
