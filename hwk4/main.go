package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Person struct {
	Firstname   string  `json:"firstname" validate:"required,lte=12,alpha,Capitalization"`
	Lastname    string  `json:"lastname" validate:"required,lte=12,alpha,Capitalization"`
	Phone       int     `json:"phone" validate:"required,lte=9999999999,gte=1111111111"`
	Email       string  `json:"email" validate:"required,email,lt=35"`
	ssn         string  `json:"ssn" validate:"required,ssn"` //show as empty lowercase field(private field not shared), or - in json tag
	MailAddress Address `json:"mailingAddress" validate:"required"`
	BillAddress Address `json:"billingAddress" validate:"required"`
}

type Address struct {
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

var ps []Person

var validate *validator.Validate

func main() {

	validate = validator.New()
	validate.RegisterValidation("Capitalization", capitalization)

	r := mux.NewRouter()
	r.HandleFunc("/healthcheck", HealthCheckHandler)
	r.HandleFunc("/test", testHandler)
	r.HandleFunc("/jsonperson", returnPerson)
	r.HandleFunc("/jsonpersons", returnPersons)
	r.HandleFunc("/jsonperson/create", createPerson)
	// http.Handle("/", r)
	srv := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:8000",
	}
	log.Fatal(srv.ListenAndServe())
}

func capitalization(fl validator.FieldLevel) bool {
	if fl.Field().String() == strings.Title(fl.Field().String()) {
		return true
	}
	return false
}

func returnPerson(w http.ResponseWriter, r *http.Request) {
	Evan := Person{
		Firstname: "Evan",
		Lastname:  "Kilpatrick",
		Phone:     7706886290,
		Email:     "ekilpatrick122495@gmail.com",
		ssn:       "985430000",
		MailAddress: Address{
			Address1: "804 Harbor Pointe Parkway",
			City:     "Roswell",
			State:    "GA",
			Zip:      "30350",
		},
		BillAddress: Address{
			Address1: "804 Harbor Pointe Parkway",
			City:     "Roswell",
			State:    "GA",
			Zip:      "30350",
		},
	}

	body, err := json.Marshal(Evan)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error-%v", err.Error())))
		return
	}

	w.Write(body)
}

func returnPersons(w http.ResponseWriter, r *http.Request) {

	body, err := json.Marshal(ps)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error-%v", err.Error())))
		return
	}

	w.Write(body)
}

func createPerson(w http.ResponseWriter, r *http.Request) {
	var p Person
	c := ps

	// var d []byte

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error-%v", err.Error())))
		return
	}

	err = validate.Struct(p)
	if err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}

		var errors APIErrors

		for _, err := range err.(validator.ValidationErrors) {

			newError := APIError{
				Field: err.Field(),
				Tag:   err.Tag(),
				Value: err.Value(),
			}

			errors.Errors = append(errors.Errors, newError)
		}

		body, _ := json.Marshal(errors)
		w.Write(body)

		return
	}
	w.Write([]byte("validation ok"))

	c = append(c, p)

	err = validate.Var(c, "unique=Email")
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}

		var errors APIErrors

		for _, err := range err.(validator.ValidationErrors) {

			newError := APIError{
				Field: err.Field(),
				Tag:   err.Tag(),
				Value: err.Value(),
			}

			errors.Errors = append(errors.Errors, newError)
		}

		body, _ := json.Marshal(errors)
		w.Write(body)

		return
	} else {
		ps = c
		w.Write([]byte("Appended"))
	}

}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("ok"))

}

func testHandler(w http.ResponseWriter, r *http.Request) {
	day := 0
	for {
		if day == 0 {
			w.Write([]byte("The 12 Days of Christmas\n\n"))
			day = day + 1
		}
		if day == 1 {
			w.Write([]byte("On the 1st day of christmas\nMy true love gave to me"))
		} else if day == 2 {
			w.Write([]byte("On the 2nd day of christmas\nMy true love gave to me"))
		} else if day == 3 {
			w.Write([]byte("On the 3rd day of christmas\nMy true love gave to me"))
		} else {
			w.Write([]byte("On the 4th day of christmas\nMy true love gave to me"))
		}

		switch day {
		case 12:
			w.Write([]byte("Twelve drummers drumming,"))
			fallthrough
		case 11:
			w.Write([]byte("Eleven pipers pipping,"))
			fallthrough
		case 10:
			w.Write([]byte("Ten lords a leaping,"))
			fallthrough
		case 9:
			w.Write([]byte("Nine ladies dancing,"))
			fallthrough
		case 8:
			w.Write([]byte("Eight maids a milking,"))
			fallthrough
		case 7:
			w.Write([]byte("Seven swans a swimming,"))
			fallthrough
		case 6:
			w.Write([]byte("Six geese-a-laying,"))
			fallthrough
		case 5:
			w.Write([]byte("Five golden rings,"))
			fallthrough
		case 4:
			w.Write([]byte("Four calling bird,"))
			fallthrough
		case 3:
			w.Write([]byte("Three french hens,"))
			fallthrough
		case 2:
			w.Write([]byte("Two turtle doves and,"))
			fallthrough
		case 1:
			w.Write([]byte("A partridge in a pear tree \n \n "))
		}
		day++
		if day > 12 {
			break
		}
	}
}
