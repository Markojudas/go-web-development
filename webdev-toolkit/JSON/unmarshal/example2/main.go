package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// the struct: note how it is incomplete!
//not using all the data just lat/long and City
type city struct {
	Latitude, Longitude float64
	City                string
}

//creating the type to store the array of the struct
type cities []city

func main() {

	var data cities

	rcvd := `[{"precision":"zip","Latitude":37.7668,"Longitude":-122.3959,"Address":"","City":"SAN FRANCISCO","State":"CA","Zip":"94107","Country":"US"},{"precision":"zip","Latitude":37.371991,"Longitude":-122.02602,"Address":"","City":"SUNNYVALE","State":"CA","Zip":"94085","Country":"US"}]`

	err := json.Unmarshal([]byte(rcvd), &data)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(data)
}
