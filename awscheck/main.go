package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type AWSKeys struct {
	AccessKeyID     string `json:"AccessKeyId"`
	SecretAccessKey string `json:"SecretAccessKey"`
	Token           string `json:"Token"`
}

func main() {
	dat, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	keys := new(AWSKeys)
	err = json.Unmarshal(dat, keys)
	if err != nil {
		log.Fatalf("error unmarshalling input, bad json provided on input: %v\n", err)
	}
	fmt.Printf("export AWS_ACCESS_KEY_ID=%s\n", keys.AccessKeyID)
	fmt.Printf("export AWS_SECRET_ACCESS_KEY=%s\n", keys.SecretAccessKey)
	fmt.Printf("export AWS_SESSION_TOKEN=%s\n", keys.Token)
}
