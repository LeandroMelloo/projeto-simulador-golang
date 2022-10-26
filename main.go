package main

import (
	"log"

	"github.com/LeandroMelloo/simulador-golang/infra/kafka"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	producer := kafka.NewKafkaProducer()
	kafka.Publish("ola", "readteste", producer)
}