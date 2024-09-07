package models

import "encoding/json"

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (p *Person) MarshalJson() ([]byte, error) {
	return json.Marshal(p)
}
