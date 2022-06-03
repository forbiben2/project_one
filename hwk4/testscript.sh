#!/bin/sh
curl -X POST 127.0.0.1:8000/jsonperson/create -H "Content-Type: application/json" -d '{"firstname":"Evan", "lastname": "Kilpatrick", "phone": 7706886290, "email": "something@gmail.com", "ssn": "890-89-0890", "mailingAddress": {"address1":"somestreet", "city": "atl", "state":"GA", "zip" : 30350}, "billingAddress":{"address1":"somestreet", "city": "atl", "state":"GA", "zip" : 30350}}'
echo
curl -X POST 127.0.0.1:8000/jsonperson/create -H "Content-Type: application/json" -d '{"lastname": "Kilpatrick", "phone": 7706886290, "email": "something@gmail.com", "ssn": "890-89-0890", "mailingAddress": {"address1":"somestreet", "city": "atl", "state":"GA", "zip" : 30350}, "billingAddress":{"address1":"somestreet", "city": "atl", "state":"GA", "zip" : 30350}}'
echo
curl -X POST 127.0.0.1:8000/jsonperson/create -H "Content-Type: application/json" -d '{"firstname":"Evan", "phone": 7706886290, "email": "something@gmail.com", "ssn": "890-89-0890", "mailingAddress": {"address1":"somestreet", "city": "atl", "state":"GA", "zip" : 30350}, "billingAddress":{"address1":"somestreet", "city": "atl", "state":"GA", "zip" : 30350}}'
echo
curl -X POST 127.0.0.1:8000/jsonperson/create -H "Content-Type: application/json" -d '{"firstname":"Evan", "lastname": "Kilpatrick", "email": "something@gmail.com", "ssn": "890-89-0890", "mailingAddress": {"address1":"somestreet", "city": "atl", "state":"GA", "zip" : 30350}, "billingAddress":{"address1":"somestreet", "city": "atl", "state":"GA", "zip" : 30350}}'
echo
curl -X POST 127.0.0.1:8000/jsonperson/create -H "Content-Type: application/json" -d '{"firstname":"Evan", "lastname": "Kilpatrick", "phone": 7706886290, "ssn": "890-89-0890", "mailingAddress": {"address1":"somestreet", "city": "atl", "state":"GA", "zip" : 30350}, "billingAddress":{"address1":"somestreet", "city": "atl", "state":"GA", "zip" : 30350}}'
echo
curl -X POST 127.0.0.1:8000/jsonperson/create -H "Content-Type: application/json" -d '{"firstname":"Evan", "lastname": "Kilpatrick", "phone": 7706886290, "email": "something@gmail.com", "ssn": "890-89-0890", "billingAddress":{"address1":"somestreet", "city": "atl", "state":"GA", "zip" : 30350}}'
echo
curl -X POST 127.0.0.1:8000/jsonperson/create -H "Content-Type: application/json" -d '{"firstname":"Evan", "lastname": "Kilpatrick", "phone": 7706886290, "email": "something@gmail.com", "ssn": "890-89-0890", "mailingAddress": {"address1":"somestreet", "city": "atl", "state":"GA", "zip" : 30350}}'