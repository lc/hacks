package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		var jsonmap = make(map[string]string)
		input := sc.Text()
		amp := strings.Split(input, "&")
		if len(amp) > 0 {
			for _, pair := range amp {
				key, val := splitter(pair)
				if key != "" {
					jsonmap[key] = val
				}
			}
		} else {
			key, val := splitter(input)
			if key != "" {
				jsonmap[key] = val
			}
		}
		out, err := json.Marshal(jsonmap)
		if err != nil {
			log.Fatalf("could not marshal map: %v", err)
		}
		fmt.Println(string(out))
	}
}
func splitter(keyval string) (string, string) {
	tmp := strings.Split(keyval, "=")
	if len(tmp) == 2 {
		return tmp[0], tmp[1]
	}
	return "", ""
}
