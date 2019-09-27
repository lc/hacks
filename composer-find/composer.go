package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type Lock struct {
	Packages []struct {
		Name    string `json:"name"`
		Version string `json:"version"`
		Source  struct {
			URL string `json:"url"`
			Ref string `json:"reference"`
		} `json:"source"`
	} `json:"packages"`
}

func main() {
	c := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		Timeout: time.Duration(10) * time.Second,
	}
	if (len(os.Args)) < 2 {
		fmt.Printf("[*] Usage: %s http://url/composer.lock\n", os.Args[0])
		os.Exit(1)
	}
	resp, err := c.Get(os.Args[1])
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	defer resp.Body.Close()
	l := Lock{}
	if err = json.NewDecoder(resp.Body).Decode(&l); err != nil {
		log.Fatalf("error decoding json: %v", err)
	}
	var url string
	for _, pkg := range l.Packages {
		url = pkg.Source.URL
		if strings.Contains(pkg.Source.URL, "https://github.com/") {
			x := strings.Split(pkg.Source.URL, ".git")
			url = fmt.Sprintf("%s/tree/%s", x[0], pkg.Source.Ref)
		}
		fmt.Printf("%s (%s) -> %s\n", pkg.Name, pkg.Version, url)
	}
}
