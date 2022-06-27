package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
)

func main() {

	code := getCode("test@example.com")
	fmt.Println(code)

	code = getCode("test@exampl.com")
	fmt.Println(code)
}

func getCode(str string) string {

	//get the hash
	hash := hmac.New(sha256.New, []byte("OURSECRET"))

	//write the string to the hash
	io.WriteString(hash, str)

	//return the hashed code (hex)

	return fmt.Sprintf("%x", hash.Sum(nil))
}
