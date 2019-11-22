package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func Readln(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}

func main() {

	domainsToMatch := make(map[string][]string)
	//categoriesToCheck := []string{"adult", "agressif", "celebrity", "dangerous_material", "dating", "drogue", "malware", "mixed_adult", "phishing", "sect", "warez"}
	categoriesToCheck := []string{"dangerous_material"}

	for _, filepath := range categoriesToCheck {
		ofd, err := os.Open("dest/" + filepath + "/domains")
		if err != nil {
			fmt.Printf("error opening file: %v\n", err)
			os.Exit(1)
		}
		reader := bufio.NewReader(ofd)
		domain, e := Readln(reader)
		for e == nil {
			domainsToMatch[filepath] = append(domainsToMatch[filepath], domain)
			domain, e = Readln(reader)
		}
	}

	results := make(map[string]string)
	re := regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)

	ofd, err := os.Open("/var/log/named/queries.log")
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}
	reader := bufio.NewReader(ofd)
	line, e := Readln(reader)
	for e == nil {
		for category, v := range domainsToMatch {
			for _, element := range v {
				if strings.Contains(line, element) {
					ip := re.FindString(line)
					results[ip + category + element] = "ip: " + ip + ", category: " + category + ", domain: " + element
				}
			}
		}
		line, e = Readln(reader)
	}
	for _, info := range results {
		fmt.Printf(info + "\n")
	}
}
