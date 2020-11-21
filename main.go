package main

import (
	"bufio"
	"fmt"
	"github.com/urfave/cli"
	"io/ioutil"
	"log"
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
	}
	return domainsToMatch
}

func printDns(){
	categories := []string{"adult", "agressif", "celebrity", "dangerous_material", "dating", "drogue", "malware", "mixed_adult", "phishing", "sect", "warez"}
	domainsToMatch := extractDomainsToMatch(categories)

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

func bind9ZonesFormat() string {
	return "zone \"%s\" {type master; file \"/etc/bind/blacklisted.db\";};\n"
}

func updateBlacklistedZones() {
	categories := []string{"agressif", "dangerous_material", "drogue", "malware", "phishing", "sect", "warez"}
	domainsToMatch := extractDomainsToMatch(categories)

	f, err := os.Create("build/blacklisted.db")
	if err != nil {
		fmt.Println(err)
		return
	}

	format := bind9ZonesFormat()
	for domain, _ := range domainsToMatch {
		f.WriteString(fmt.Sprintf(format, domain))
	}

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(fmt.Sprintf("%d domains added to blacklist", len(domainsToMatch)))
}

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "show",
				Aliases: []string{"s"},
				Usage:   "show dns already queried",
				Action: func(c *cli.Context) error {
					printDns()
					return nil
				},
			},
			{
				Name:    "update",
				Aliases: []string{"u"},
				Usage:   "update blacklisted domains",
				Action: func(c *cli.Context) error {
					updateBlacklistedZones()
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
