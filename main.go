package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	loadEnvironmentVariable()

	log.Println("Hello World!")
}

func loadEnvironmentVariable() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}
}
