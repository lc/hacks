package main

import (
	"crypto/tls"
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

type WadlApp struct {
	Resources struct {
		Text     string `xml:",chardata"`
		Base     string `xml:"base,attr"`
		Resource []struct {
			Text     string `xml:",chardata"`
			Path     string `xml:"path,attr"`
			Resource []struct {
				Text  string `xml:",chardata"`
				Path  string `xml:"path,attr"`
				Param []struct {
					Text  string `xml:",chardata"`
					Xs    string `xml:"xs,attr"`
					Name  string `xml:"name,attr"`
					Style string `xml:"style,attr"`
					Type  string `xml:"type,attr"`
				} `xml:"param"`
				Method struct {
					Text string `xml:",chardata"`
					ID   string `xml:"id,attr"`
					Name string `xml:"name,attr"`
				} `xml:"method"`
			} `xml:"resource"`
		} `xml:"resource"`
	} `xml:"resources"`
}

func main() {
	url := flag.String("u", "", "url of application.wadl file")
	paramsOnly := flag.Bool("p", false, "dump all params only")
	flag.Parse()

	if *url == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	cl := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			DialContext: (&net.Dialer{
				Timeout: 10 * time.Second,
			}).DialContext,
		},
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	resp, err := cl.Get(*url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var wadl WadlApp
	err = xml.Unmarshal(body, &wadl)
	if err != nil {
		log.Fatal(err)
	}
	resources := wadl.Resources
	apiBase := resources.Base
	if *paramsOnly {
		for _, re := range resources.Resource {
			for _, me := range re.Resource {
				for _, fs := range me.Param {
					fmt.Printf("%s\n", fs.Name)
				}
			}
		}
	} else {
		for _, re := range resources.Resource {
			base := fmt.Sprintf("%s%s", apiBase, strings.TrimPrefix(re.Path, "/"))
			for _, me := range re.Resource {
				fmt.Printf("%s %s%s\n", me.Method.Name, base, me.Path)
			}
		}
	}
}
