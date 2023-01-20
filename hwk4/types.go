package main

import "gorm.io/gorm"

type Person struct {
	gorm.Model
	Firstname   string  `json:"firstname" validate:"required,lte=12,alpha,Capitalization"`
	Lastname    string  `json:"lastname" validate:"required,lte=12,alpha,Capitalization"`
	Phone       int     `json:"phone" validate:"required,lte=9999999999,gte=1111111111"`
	Email       string  `json:"email" validate:"required,email,lt=35"`
	ssn         string  `json:"ssn" validate:"required,ssn"` //show as empty lowercase field(private field not shared), or - in json tag
	MailAddress Address `json:"mailingAddress" validate:"required"`
	BillAddress Address `json:"billingAddress" validate:"required"`
}

type Address struct {
	gorm.Model
	Address1 string `json:"address1" validate:"required,lt=30"`
	Address2 string `json:"address2,omitempty" validate:"lt=30"` //will omit if empty
	City     string `json:"city" validate:"required,lt=15,alpha,Capitalization"`
	State    string `json:"state" validate:"required,len=2,alpha,uppercase"`
	Zip      string `json:"zip" validate:"required,numeric,len=5"`
}

type APIErrors struct {
	Errors []APIError
}
type APIError struct {
	Field string
	Tag   string
	Value interface{}
}
