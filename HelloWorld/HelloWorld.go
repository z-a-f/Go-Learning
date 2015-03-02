package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed( time.Now().UTC().UnixNano())
	r := rand.Intn(100)
	fmt.Printf ("%d. HelloWorld!\n", r)

}
