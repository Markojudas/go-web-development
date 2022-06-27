package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {

	str := "For a decent price I've banned kindness from my heart. The spirit of all truth and beauty pawned for my desire"

	encodeStd := "ABCDEFGHIJKLMNOPQRStUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

	s64 := base64.NewEncoding(encodeStd).EncodeToString([]byte(str))

	//or if you don't want to specify the endocing standard:
	// s64 := base64.StdEncoding.EncodeToString([]byte(str))

	fmt.Println(len(str))
	fmt.Println(len(s64))
	fmt.Println(str)
	fmt.Println(s64)

	//decoding
	bs, err := base64.NewEncoding(encodeStd).DecodeString(s64)
	if err != nil {
		log.Fatalln("ERROR!!!")
	}

	fmt.Println(string(bs))
}
