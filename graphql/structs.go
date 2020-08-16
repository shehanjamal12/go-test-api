package main

type Items struct {
	ItemName     string `json:"itemName"` //json part so that in postman it comes as simple i not capital
	ItemQuantity int    `json:"itemQuantity"`
}