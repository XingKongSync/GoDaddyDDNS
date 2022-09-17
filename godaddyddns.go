package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

var (
	HOST    string
	DOMAIN  string
	SECRET  string
	KEY     string
	API_URL = "https://api.godaddy.com/v1/domains/%s/records/%s/%s"
)

const TYPE = "A"

type DNSRecord struct {
	Data string `json:"data"`
	Name string `json:"name"`
	Type string `json:"type"`
}

func handleErr(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
}

func initFlags() {
	flag.StringVar(&HOST, "host", "", "your host")
	flag.StringVar(&DOMAIN, "domain", "", "your domain")
	flag.StringVar(&SECRET, "secret", "", "your api secret")
	flag.StringVar(&KEY, "key", "", "your api key")

	flag.Parse()

	// fmt.Printf("Host: %s\n", HOST)
	// fmt.Printf("Domain: %s\n", DOMAIN)
	// fmt.Printf("Key: %s\n", KEY)
	// fmt.Printf("Secret: %s\n", SECRET)

	API_URL = fmt.Sprintf(API_URL, DOMAIN, TYPE, HOST)

	// fmt.Printf("API URL: %s\n", API_URL)
}

func initHeader(header *http.Header) {
	header.Add("Accept", "application/json")
	header.Add("Content-type", "application/json")
	header.Add("Authorization", fmt.Sprintf("sso-key %s:%s", KEY, SECRET))
}

func getDNS() string {
	req, err := http.NewRequest("GET", API_URL, nil)
	handleErr(err)
	initHeader(&req.Header)

	// fmt.Println(req.Header)

	var record []DNSRecord
	resp, err := http.DefaultClient.Do(req)

	handleErr(err)

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	handleErr(err)

	// fmt.Printf("%s\n", body)
	err = json.Unmarshal(body, &record)
	handleErr(err)

	if len(record) <= 0 {
		return ""
	}
	return record[0].Data
}

func getIP() string {
	const CHECK_URL = "http://api.ipify.org"

	resp, err := http.Get(CHECK_URL)
	handleErr(err)

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	handleErr(err)

	return string(body)
}

func setDNS(ip string) {

	reqBody := []DNSRecord{
		{Data: ip, Type: TYPE, Name: HOST},
	}

	bytes, err := json.Marshal(reqBody)
	handleErr(err)

	req, err := http.NewRequest("PUT", API_URL, strings.NewReader(string(bytes)))
	handleErr(err)
	initHeader(&req.Header)

	resp, err := http.DefaultClient.Do(req)
	resp.Body.Close()
	handleErr(err)
}

func main() {
	fmt.Println("Preparing for ddns...")

	initFlags()

	ip := getIP()
	fmt.Printf("Current ip address is: %s\n", ip)

	dns := getDNS()
	fmt.Printf("Current record address is: %s\n", dns)

	if ip != dns {
		setDNS(ip)
		fmt.Println("Update dns successfully!")
	}
}
