package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonArr(t *testing.T)  {
	customer := Customer {
		FirstName: "Rahasia",
		MiddleName: "Abrar",
		LastName: "Naim",
		Age: 16,
		Married: false,
		Hobbies: []string {"Gaming", "Reading", "Coding"},
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
	// {"FirstName":"Rahasia","MiddleName":"Abrar","LastName":"Naim","Age":16,"Married":false,"Hobbies":["Gaming","Reading","Coding"]}
}

func TestJsonArrdecode(t *testing.T)  {
	jsonStr := `{"FirstName":"Rahasia","MiddleName":"Abrar","LastName":"Naim","Age":16,"Married":false,"Hobbies":["Gaming","Reading","Coding"]}`
	jsonBytes := []byte(jsonStr)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer) // &{Rahasia Abrar Naim 16 false [Gaming Reading Coding]}
	fmt.Println(customer.Hobbies) // [Gaming Reading Coding]
}

func TestJsonArrComplexEncode(t *testing.T)  {
	customer := Customer {
		FirstName: "Rahasia",
		MiddleName: "Abrar",
		LastName: "Naim",
		Age: 16,
		Married: false,
		Hobbies: []string {"Gaming", "Reading", "Coding"},
		Address: []Address {
			{
				Street: "jl. Tubagus Rangin",
				Country: "Indonesia",
				PostalCode: "1111",
			},
			{
				Street: "jl. Rajin",
				Country: "Indonesia",
				PostalCode: "1111",
			},
		},
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestJsonArrComplexDecode(t *testing.T)  {
	jsonStr := `{"FirstName":"Rahasia","MiddleName":"Abrar","LastName":"Naim","Age":16,"Married":false,"Hobbies":["Gaming","Reading","Coding"],"Address":[{"Street":"jl. Tubagus Rangin","Country":"Indonesia","PostalCode":"1111"},{"Street":"jl. Rajin","Country":"Indonesia","PostalCode":"1111"}]}`
	jsonBytes := []byte(jsonStr)

	customer := &Customer {}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.Address) // [{jl. Tubagus Rangin Indonesia 1111} {jl. Rajin Indonesia 1111}]
}

func TestOnlyJsonArrComplexDecode(t *testing.T)  {
	jsonStr := `[{"Street":"jl. Tubagus Rangin","Country":"Indonesia","PostalCode":"1111"},{"Street":"jl. Rajin","Country":"Indonesia","PostalCode":"1111"}]`
	jsonBytes := []byte(jsonStr)

	address := &[]Address{}

	err := json.Unmarshal(jsonBytes, address)
	if err != nil {
		panic(err)
	}

	fmt.Println(address) // &[{jl. Tubagus Rangin Indonesia 1111} {jl. Rajin Indonesia 1111}]
}

func TestOnlyJsonArrComplexEncode(t *testing.T)  {
		address := []Address {
			{
				Street: "jl. Tubagus Rangin",
				Country: "Indonesia",
				PostalCode: "1111",
			},
			{
				Street: "jl. Rajin",
				Country: "Indonesia",
				PostalCode: "1111",
			},
		}


	bytes, err := json.Marshal(address)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
	// [{"Street":"jl. Tubagus Rangin","Country":"Indonesia","PostalCode":"1111"},{"Street":"jl. Rajin","Country":"Indonesia","PostalCode":"1111"}]
}