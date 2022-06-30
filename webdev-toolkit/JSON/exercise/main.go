package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// struct gotten from https://mholt.github.io/json-to-go/
// copy and pasted the JSON (as seen on the HTML file)

type codes []struct {
	Code    int    `json:"Code"`
	Descrip string `json:"Descrip"`
}

func main() {

	//created the variable of type codes []struct
	var data codes

	rcvd := `[{"Code":200,"Descrip":"StatusOK"},{"Code":301,"Descrip":"StatusMovedPermanently"},{"Code":302,"Descrip":"StatusFound"},{"Code":303,"Descrip":"StatusSeeOther"},{"Code":307,"Descrip":"StatusTemporaryRedirect"},{"Code":400,"Descrip":"StatusBadRequest"},{"Code":401,"Descrip":"StatusUnauthorized"},{"Code":402,"Descrip":"StatusPaymentRequired"},{"Code":403,"Descrip":"StatusForbidden"},{"Code":404,"Descrip":"StatusNotFound"},{"Code":405,"Descrip":"StatusMethodNotAllowed"},{"Code":418,"Descrip":"StatusTeapot"},{"Code":500,"Descrip":"StatusInternalServerError"}]`

	// passed the JSON as slice of bytes and the address of the 'data' variable
	//to unmarshal
	err := json.Unmarshal([]byte(rcvd), &data)
	if err != nil {
		log.Fatalln(err)
	}

	//ranging over 'data'
	for _, val := range data {
		fmt.Println(val.Code, ": ", val.Descrip)
	}
}
