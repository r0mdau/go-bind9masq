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
		isPrefix  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}

func extractDomainsFromLogs() [][]string {
	// storing entire file as string to find in one regex every domain + ip
	b, err := ioutil.ReadFile(config.Bind9.Queries)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	str := string(b)
	re := regexp.MustCompile(`((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3})#.+\((([a-z]+\.)+[a-z]{2,36})\)`)
	matches := re.FindAllStringSubmatch(str, -1)
	return matches
}

func extractDomainsToMatch(categories []string) map[string]string {
	domainsToMatch := make(map[string]string)
	// Put domains as key in map for faster finding and their category as value
	for _, category := range categories {
		ofd, err := os.Open("dest/" + category + "/domains")
		if err != nil {
			fmt.Printf("error opening file: %v\n", err)
			os.Exit(1)
		}
		reader := bufio.NewReader(ofd)
		domain, e := Readln(reader)
		for e == nil {
			domainsToMatch[domain] = category
			domain, e = Readln(reader)
		}
		ofd.Close()
	}
	return domainsToMatch
}

func writeFileFromStringSlices(domainsToMatch map[string]string) {
	f, err := os.Create("build/blacklisted.zones")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	format := bind9ZonesFormat()
	for domain, _ := range domainsToMatch {
		f.WriteString(fmt.Sprintf(format, domain))
	}

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}