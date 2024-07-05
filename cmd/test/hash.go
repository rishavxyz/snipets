package test

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"time"
)

func Hash() {
	var (
		orig_password = "rishav"
		hash_password = hash(&orig_password)
	)
	fmt.Printf("HASHED %s", hash_password)
}

func hash(password *string) string {
	hash := sha512.New()
	hash.Write([]byte(*password))
	return hex.EncodeToString(hash.Sum(nil))
}

func Show() {
	for i := 0; i < 10; i++ {
		s := time.Now()
		Hash()
		e := time.Now()
		fmt.Printf("\nTook %v\n", e.Sub(s))
	}
}
