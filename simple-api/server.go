package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// serialization: go obj{} (instance of a struct) -> json-string ([]byte)
// deserialization: the opposite
// json.marshal()/json.unmarshal() - for in-memory json processing/ops (ok for small-mid sized dataset)
// json.encoder/decoder - for streaming json data.. working with large data sets and networking connections (more flexible and suited for complex ops.)

type User struct {
	Name string `json:"name"`
	Email string `json:"email"`
}

func main() {
	user:= User{Name: "Skyy", Email: "skyy@email.com"}

	// MARSHAL
	jsonData,err:=json.Marshal(user)
	if err!=nil{
		log.Fatal("⚠️ ERR:",err)
	}
	fmt.Println(user) // {Skyy skyy@email.com}
	fmt.Println(string(jsonData)) // {"name":"Skyy","email":"skyy@email.com"}

	// UNMARSHAL
	var user1 User
	err = json.Unmarshal(jsonData,&user1)
	if err!=nil{
		log.Fatal("⚠️ ERR:",err)
	}
	fmt.Println("User created from json data:",user1) // User created from json data: {Skyy skyy@email.com}

	// ENCODER & DECODER - More common, dealing with APIs
	jsonData1:= `{"name":"Soumadip","email":"soumadip@email.com"}`
	reader:= strings.NewReader(jsonData1)
	decoder:= json.NewDecoder(reader)
	var user2 User
	err = decoder.Decode(&user2)
	if err!=nil{
		log.Fatal("⚠️ ERR:",err)
	}
	fmt.Println(user2) // {Soumadip soumadip@email.com}

	
	var buff bytes.Buffer
	encoder := json.NewEncoder(&buff) // encoder needs buffer

	err=encoder.Encode(user)
	if err!=nil{
		log.Fatal("⚠️ ERR:",err)
	}
	fmt.Println("Encoded json-string:",buff.String()) // Encoded json-string: {"name":"Skyy","email":"skyy@email.com"}
}