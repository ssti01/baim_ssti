package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func main() {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	fmt.Println(hex.EncodeToString(b))
}
