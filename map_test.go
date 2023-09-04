package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapDecode(t *testing.T)  {
	jsonStr := `{"id":"A001", "name":"Apple Mac Air M1", "price": 200000}`
	jsonByets := []byte(jsonStr)

	var result map[string] interface {}

	err := json.Unmarshal(jsonByets, &result)
	if err != nil {
		panic(err)
	}

	fmt.Println(result) // map[id:A001 name:Apple Mac Air M1 price: 200000]
	fmt.Println(result["id"]) // A001
	fmt.Println(result["name"]) // Apple Mac Air M1
	fmt.Println(result["price"]) // 200000
}

func TestMapEncode(t *testing.T)  {
	product := map[string] interface {} {
		"id": "A001",
		"name": "Apple Mac Ari M1",
		"price": 20000,
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}