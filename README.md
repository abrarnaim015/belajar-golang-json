# Belajar-Golang-JSON

Link Pembelajaran <a href="https://www.youtube.com/watch?v=oYJaWttamzE&list=PL-CtdCApEFH-0i9dzMzLw6FKVrFWv3QvQ&index=12&ab_channel=ProgrammerZamanNow">Programmer Zaman Now</a>

Link Documentasi <a href="https://www.json.org/json-en.html">JSON</a>

> Encode JSON

- Function untuk melakukan konversi data ke `JSON` dapat menggunakan function `json.Marshal(interface {})`
- Karena parameternya adalah `interface {}` maka kita bisa memasukan type data apapun ke dalam `function Marshal`

```golang
func logJson(data interface {})  {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestEncode(t *testing.T)  {
	logJson("Naim")
	logJson(1)
	logJson(true)
	logJson([]string {"Abrar", "Naim"})
}

// Result
// "Naim"
// 1
// true
// ["Abrar","Naim"]
```

> Json Object

```golang
type Customer struct {
	FirstName string
	MiddleName string
	LastName string
	Age int
	Married bool
}

func TestJsonObject(t *testing.T)  {
	customer := Customer {
		FirstName: "Rahasia",
		MiddleName: "Abrar",
		LastName: "Naim",
		Age: 26,
		Married: false,
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
	// {"FirstName":"Rahasia","MiddleName":"Abrar","LastName":"Naim","Age":26,"Married":false}
}
```

> Decode Json

- Untuk melakukan Konversi dari `JSON` ke type data di golang, kuta bisa menggunakan function `json.Unmarshal(byte[], interface {})`
- `byte[]` adalah data `JSON`nya
- `interface {}` adalah tempat menyimpan hasil konversi, biasa berupa `pointer`

```golang
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
```

> Json Arr Complex

```golang
type Address struct {
	Street string
	Country string
	PostalCode string
}

type Customer struct {
	FirstName string
	MiddleName string
	LastName string
	Age int
	Married bool
	Hobbies []string
	Address []Address
}
```

```golang
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
```

> Only Json Arr

```golang
func TestOnlyJsonArrComplexDecode(t *testing.T)  {
	jsonStr := `[{"Street":"jl. Tubagus Rangin","Country":"Indonesia","PostalCode":"1111"},{"Street":"jl. Rajin","Country":"Indonesia","PostalCode":"1111"}]`
	jsonBytes := []byte(jsonStr)

	address := &[]Address{}

	err := json.Unmarshal(jsonBytes, address)
	if err != nil {
		panic(err)
	}

	fmt.Println(address)
  // &[{jl. Tubagus Rangin Indonesia 1111} {jl. Rajin Indonesia 1111}]
}
```

```golang
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
```

> Json Tag

- kadang ada style yang berbeda antara penamaan atribut di `Struct` dan di `JSON`, misal di `JSON` kita ingin menggunakan `snake_case`, tp di `Struct`, kita ingin menggunakan `PascalCase`
- `package json` menduku `Tag Reflaction`
- kita bisa menambakan `tag reflagtin` dgn nama jsonm lalu diikuti dengan atribut yang kita inginkan ketika di konversi dari atau ke `JSON`

```golang
type Product struct {
  Id string `json: "id"`
  Name string `json: "name"`
  Price int64 `json: "price"`
  ImageURL string `json: "image_url"`
}
```

```golang
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
```

> Map Json

```golang
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
```

```golang
func TestMapEncode(t *testing.T)  {
	product := map[string] interface {} {
		"id": "A001",
		"name": "Apple Mac Ari M1",
		"price": 20000,
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}
```

> Streaming Decoder Json

- Untuk membuat `Json Decoder`, kita bisa menggunakan function `json.NewDecoder(reader)`
- Selanjutnya untuk membaca isi input `reader` dan konversikan secara langsung ke data di golang cukup menggunakan functin `Decode(interface{})`

```json
// Customer.json
{
  "FirstName": "Rahasia",
  "MiddleName": "Abrar",
  "LastName": "Naim"
}
```

```golang
func TestStreamDecoder(t *testing.T)  {
	reader, _ := os.Open("Customer.json")
	decoder := json.NewDecoder(reader)

	customer := &Customer{}
	decoder.Decode(customer)

	fmt.Println(customer)
	// &{Rahasia Abrar Naim 0 false [] []}
}
```

> Streaming Encoder json

- untuk membuat `Encoder` kita bisa menggunakan function `json.NewEncoder(writer)`
- dan untuk menulis data sebagai JSON langsung ke `writer`, kita bisa gunakan function `Encode(interface{})`

```golang
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
```

```json
// test_encoder.json
{
  "FirstName": "Rahasia",
  "MiddleName": "Abrar",
  "LastName": "Naim",
  "Age": 0,
  "Married": false,
  "Hobbies": null,
  "Address": null
}
```
