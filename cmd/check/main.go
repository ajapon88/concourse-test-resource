package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/ajapon88/concourse-test-resource"
)

type Request struct {
	Source  resource.Source  `json:"source"`
	Version resource.Version `json:"version"`
}

type Response []resource.Version

func main() {
	var request Request
	decoder := json.NewDecoder(os.Stdin)
	err := decoder.Decode(&request)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to decode: %s\n", err.Error())
		os.Exit(1)
		return
	}

	fmt.Fprintf(os.Stderr, "source: %v\n", request.Source)

	response := Response{}
	response = append(response, resource.Version{Date: time.Now().Format("2006-01-02 15:04")})

	json.NewEncoder(os.Stdout).Encode(response)
}
