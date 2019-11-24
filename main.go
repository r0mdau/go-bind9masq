package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
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

	domainsToMatch := make(map[string]string)
	categoriesToCheck := []string{"adult", "agressif", "celebrity", "dangerous_material", "dating", "drogue", "malware", "mixed_adult", "phishing", "sect", "warez"}

	// Put domains as key in map for faster finding and their category as value
	for _, filepath := range categoriesToCheck {
		ofd, err := os.Open("dest/" + filepath + "/domains")
		if err != nil {
			fmt.Printf("error opening file: %v\n", err)
			os.Exit(1)
		}
		reader := bufio.NewReader(ofd)
		domain, e := Readln(reader)
		for e == nil {
			domainsToMatch[domain] = filepath
			domain, e = Readln(reader)
		}
	}

	// storing entire file as string to find in one regex every domain + ip
	b, err := ioutil.ReadFile("/var/log/named/queries.log")
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)
	re := regexp.MustCompile(`((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3})#.+\((([a-z]+\.)+[a-z]{2,36})\)`)
	matches := re.FindAllStringSubmatch(str, -1)

	// printable map to avoid duplicate
	printableResults := make(map[string]string)
	for _, results := range matches {
		if _, ok := domainsToMatch[results[5]]; ok {
			category := domainsToMatch[results[5]]
			printableResults[results[1]+category+results[5]] = "ip: " + results[1] + ", category: " + category + ", domain: " + results[5]
		}
	}

	for _, info := range printableResults {
		fmt.Printf(info + "\n")
	}
}
