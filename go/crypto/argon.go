package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"

	"golang.org/x/crypto/argon2"
)

func main() {
	plain := "password123"

	// Salt should be random, hence the crypto/rand function of the rand.Reader
	salt := make([]byte, 32)
	io.ReadFull(rand.Reader, salt)

	key := argon2.IDKey([]byte(plain), salt, 1, 64*1024, 4, 64)
	fmt.Println(key)
	// Key is type of []byte, to make it readable, convert it to string via hex.EncodeToString
	fmt.Println(hex.EncodeToString(key))
}
