package main

// uwords.go by corben leo (@hacker_)
// gets all unique words given a host
// and a list of paths

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strings"
	"time"
)

func main() {
	data := &bytes.Buffer{}
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s url\n", os.Args[0])
		os.Exit(1)
	}
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		Timeout: time.Duration(10) * time.Second,
	}
	tg := os.Args[1]
	resp, err := client.Get(tg)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading body: %v", err)
	}
	data.Write(body)

	words := regexp.MustCompile(`[^a-zA-Z0-9_-]`).Split(data.String(), -1)
	if words != nil {
		sort.Strings(words)
		res := dupe(words)
		for _, word := range res {
			word = strings.Replace(word, " ", "", -1)
			word = strings.Replace(word, "\n", "", -1)
			word = strings.Replace(word, "--", "", -1)
			word = strings.TrimPrefix(word, "-")
			word = strings.TrimSpace(strings.Replace(word, " ", " ", -1))
			if word != "" {
				fmt.Println(word)
			}
		}
	}
}

// taken from subfinder
// libsubfinder/helper/misc.go#L51
func dupe(elements []string) []string {
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{}
	result := []string{}
	for v := range elements {
		if encountered[elements[v]] {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}
