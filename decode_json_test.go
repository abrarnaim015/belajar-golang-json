package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJson(t *testing.T)  {
	jsonString := `{"FirstName":"Rahasia","MiddleName":"Abrar","LastName":"Naim","Age":26,"Married":false}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer) // &{Rahasia Abrar Naim 26 false}
	fmt.Println(customer.FirstName) // Rahasia
	fmt.Println(customer.MiddleName) // Abrar
	fmt.Println(customer.LastName) // Naim
}