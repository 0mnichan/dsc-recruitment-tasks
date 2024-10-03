package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func hasher() {
	var pwd string
	fmt.Println("pls enter your pwd")
	fmt.Scanln(&pwd)
	hash := sha256.New()
	hash.Write([]byte(pwd))
	hashed := hex.EncodeToString(hash.Sum(nil))
	fmt.Println("SHA-256 hash of your pwd is:", hashed)
}

func main() {
	hasher()
}
