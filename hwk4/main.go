package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	Firstname   string  `json:"firstname"`
	Lastname    string  `json:"lastname"`
	Phone       int     `json:"phone"`
	Email       string  `json:"email"`
	ssn         string  `json:"ssn"` //show as empty lowercase field(private field not shared), or - in json tag
	MailAddress Address `json:"mailingAddress"`
	BillAddress Address `json:"billingAddress"`
}

type Address struct {
	Address1 string `json:"address1"`
	Address2 string `json:"address2,omitempty"` //will omit if empty
	City     string `json:"city"`
	State    string `json:"state"`
	Zip      int    `json:"zip"`
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/healthcheck", HealthCheckHandler)
	r.HandleFunc("/test", testHandler)
	r.HandleFunc("/jsonperson", returnPerson)
	r.HandleFunc("/jsonperson/create", createPerson)
	// http.Handle("/", r)
	srv := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:8000",
	}
	log.Fatal(srv.ListenAndServe())
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
			Zip:      30350,
		},
		BillAddress: Address{
			Address1: "804 Harbor Pointe Parkway",
			City:     "Roswell",
			State:    "GA",
			Zip:      30350,
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

func createPerson(w http.ResponseWriter, r *http.Request) {
	var p Person
	// var d []byte

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error-%v", err.Error())))
		return
	}

	// json.Unmarshal(d, &p)
	if p.Firstname == "" {
		w.Write([]byte("One or more fields missing(firstname)"))
	} else if p.Lastname == "" {
		w.Write([]byte("One or more fields missing(lastname)"))
	} else if p.Phone == 0 {
		w.Write([]byte("One or more fields missing(phone)"))
	} else if p.Email == "" {
		w.Write([]byte("One or more fields missing(email)"))
	} else if p.MailAddress.Address1 == "" {
		w.Write([]byte("One or more fields missing(address1 mailaddress)"))
	} else if p.MailAddress.City == "" {
		w.Write([]byte("One or more fields missing(city mailaddress)"))
	} else if p.MailAddress.State == "" {
		w.Write([]byte("One or more fields missing(state mailaddress)"))
	} else if p.MailAddress.Zip == 0 {
		w.Write([]byte("One or more fields missing(zip mailaddress)"))
	} else if p.BillAddress.Address1 == "" {
		w.Write([]byte("One or more fields missing(address1 billaddress)"))
	} else if p.BillAddress.City == "" {
		w.Write([]byte("One or more fields missing(city billaddress)"))
	} else if p.BillAddress.State == "" {
		w.Write([]byte("One or more fields missing(state billaddress)"))
	} else if p.BillAddress.Zip == 0 {
		w.Write([]byte("One or more fields missing(zip billaddress)"))
	} else {
		w.Write([]byte(fmt.Sprintf("%+v", p)))
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
