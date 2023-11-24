package model

import "encoding/json"

type Cart struct {
	Id    string   `json:"id"`
	Items []string `json:"items"`
}

func DecodeCart(rawCart string) (*Cart, error) {
	cart := &Cart{}
	err := json.Unmarshal([]byte(rawCart), cart)
	return cart, err
}

func EncodeCart(cart *Cart) (string, error) {
	rawCart, err := json.Marshal(cart)
	return string(rawCart), err
}
