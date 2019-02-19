package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

func httpPostJSON(url string, v interface{}, res interface{}) (*http.Response, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(serverAddr+url, "application/json", bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if res != nil {
		b, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return resp, err
		}
		if err := json.Unmarshal(b, res); err != nil {
			fmt.Println("Err resp body:", string(b))
			return resp, err
		}
		return resp, nil
	}
	return resp, nil
}

func httpGetJSON(url string, res interface{}) (*http.Response, error) {
	resp, err := http.Get(serverAddr + url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if res != nil {
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return resp, err
		}
		if err := json.Unmarshal(b, res); err != nil {
			fmt.Println("Err resp body:", string(b))
			return resp, err
		}
		return resp, nil
	}
	return resp, nil
}
