package main

import (
	"flag"
	"fmt"

	"gopkg.in/hlandau/passlib.v1"
)

func main() {
	flag.Parse()
	tail := flag.Args()

	if len(tail) == 1 {
		// hashing password
		hash, _ := passlib.Hash(tail[0])
		fmt.Println(hash)
	}
	if len(tail) == 2 {
		// verifying password
		_, err := passlib.Verify(tail[0], tail[1])

		if err != nil {
			fmt.Println(0)
			fmt.Println(err)
		} else {
			fmt.Println(1)
		}
	}
}
