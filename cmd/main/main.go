package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	failOnError(err, fmt.Sprintf("Error loading .env file"), false)

	// env variable extraction
	EnvAMPQQueryString := os.Getenv("AMPQ_QUERYSTRING")
	EnvQueueName := os.Getenv("QUEUE_NAME")
	EnvCmd := os.Getenv("CMD")
	EnvInterval, err := strconv.Atoi(os.Getenv("INTERVAL"))
	failOnError(err, fmt.Sprintf("cannot convert %s to int", os.Getenv("INTERVAL")), false)
	failOnError(err, fmt.Sprintf("cannot convert %s to int", os.Getenv("PAYLOAD_SIZE")), false)
	if EnvInterval == 0 {
		// set default value in error case
		EnvInterval = 1000
	}
	EnvPayloadSize, err := strconv.Atoi(os.Getenv("PAYLOAD_SIZE"))
	failOnError(err, fmt.Sprintf("cannot convert %s to int", os.Getenv("PAYLOAD_SIZE")), false)
	if EnvPayloadSize == 0 {
		// set default value in error case
		EnvPayloadSize = 128
	}

	// using standard library "flag" package
	cmd := flag.String("cmd", EnvCmd, "cmd to send")
	interval := flag.Int("interval", EnvInterval, "time window to call inner cmd")
	payloadSize := flag.Int("size", EnvPayloadSize, "time window to call inner cmd")
	flag.Parse()

	log.Printf("running %v cmd", cmd)
	if *cmd == "send" {
		send(EnvAMPQQueryString, EnvQueueName, *interval, *payloadSize)
	} else {
		receive(EnvAMPQQueryString, EnvQueueName, *interval)
	}
}
