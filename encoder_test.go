package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamEncoder(t *testing.T)  {
	writer, err := os.Create("test_encoder.json")
	if err != nil {
		panic(err)
	}
	encoder := json.NewEncoder(writer)

	customer := Customer {
		FirstName: "Rahasia",
		MiddleName: "Abrar",
		LastName: "Naim",
	}

	err = encoder.Encode(customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
}