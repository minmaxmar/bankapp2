package main

import (
	"bankapp2/bootstrapper"
	"log"
)

func main() {
	err := bootstrapper.New().RunAPI()
	if err != nil {
		log.Fatal("failed to start")
	}
}
