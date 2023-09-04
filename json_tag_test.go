package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
  Id string `json:"id"`
  Name string `json:"name"`
  Price int64 `json:"price"`
  ImageURL string `json:"image_url"`
}

func TestJsonTag(t *testing.T)  {
	product := Product {
		Id: "A001",
		Name: "Apple Mac Air M1",
		ImageURL: "http://test.com/image.png",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
	// {"id":"A001","name":"Apple Mac Air M1","price":0,"image_url":"http://test.com/image.png"}
}

func TestJsonTagDecode(t *testing.T)  {
	jsonStr := `{"id":"A001","name":"Apple Mac Air M1","price":0,"image_url":"http://test.com/image.png"}`
	jsonBytes := []byte(jsonStr)

	product := &Product {}

	err := json.Unmarshal(jsonBytes, product)
	if err != nil {
		panic(err)
	}
	fmt.Println(product)
}