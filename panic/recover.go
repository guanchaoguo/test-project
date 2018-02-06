package main

import (
	"os"
	"log"
)
var user = os.Getenv("USER")
func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("work failed:", err)
		}
	}()

	if user == "" {
		panic("no value for $USER")
	}


}

