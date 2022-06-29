package main

import (
	"encoding/json"
	"fmt"
	"log"
)

//broke down the struct into 2 structs

type thumbnail struct {
	URL           string
	Height, Width int
}

type img struct {
	Width, Height int
	Title         string
	Thumbnail     thumbnail //here's the structure!!!
	Animated      bool
	IDs           []int
}

func main() {

	var data img

	rcvd := `{"Width":800,"Height":600,"Title":"View from 15th Floor","Thumbnail":{"Url":"http://www.example.com/image/481989943","Height":125,"Width":100},"Animated":false,"IDs":[116,943,234,38793]}`

	err := json.Unmarshal([]byte(rcvd), &data)
	if err != nil {
		log.Fatalln("error unmarshelling", err)
	}

	fmt.Println(data)

	for idx, val := range data.IDs {
		fmt.Println(idx, val)
	}

	fmt.Println(data.Thumbnail.URL)

}
