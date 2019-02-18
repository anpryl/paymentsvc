package main

import (
	"log"
	"os"
	"testing"
)

var serverAddr string

func TestMain(m *testing.M) {
	v := os.Getenv("INTEGRATION_TESTING")
	if v == "" {
		log.Println("Skipping integration testing")
		return
	}
	serverAddr = os.Getenv("TEST_SERVER_ADDR")
	if serverAddr == "" {
		log.Fatalln("Empty test server address")
		return
	}
	os.Exit(m.Run())
}
