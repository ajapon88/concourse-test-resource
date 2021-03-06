package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/ajapon88/concourse-test-resource"
)

type Request struct {
	Source resource.Source `json:"source"`
	Params Params          `json:"params"`
}

type Params struct {
	PutParam1 string `json:"put_param1"`
}

type Response struct {
	Version  resource.Version        `json:"version"`
	Metadata []resource.MetadataPair `json:"metadata"`
}

func main() {
	var request Request
	decoder := json.NewDecoder(os.Stdin)
	err := decoder.Decode(&request)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to decode: %s\n", err.Error())
		os.Exit(1)
		return
	}

	// 今回はresourceに対する操作がないためsrcは必要ない。
	//src := os.Args[1]
	fmt.Fprintf(os.Stderr, "source: %v\n", request.Source)
	fmt.Fprintf(os.Stderr, "params: %v\n", request.Params)

	t := time.Now()
	response := Response{
		resource.Version{Date: t.Format("2006-01-02 15:04")},
		[]resource.MetadataPair{
			{Name: "Year", Value: strconv.Itoa(t.Year())},
			{Name: "Month", Value: t.Month().String()},
			{Name: "Day", Value: strconv.Itoa(t.Day())},
		},
	}

	json.NewEncoder(os.Stdout).Encode(response)
}
