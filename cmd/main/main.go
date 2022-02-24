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
	if EnvInterval == 0 {
		// set default value in error case
		EnvInterval = 1000
	}

	// using standard library "flag" package
	cmd := flag.String("cmd", EnvCmd, "cmd to send")
	interval := flag.Int("interval", EnvInterval, "time window to call inner cmd")
	flag.Parse()

	log.Printf("running %v cmd", cmd)
	if *cmd == "send" {
		send(EnvAMPQQueryString, EnvQueueName, *interval)
	} else {
		receive(EnvAMPQQueryString, EnvQueueName, *interval)
	}
}
