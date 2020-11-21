package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

func deduplicateResults(matches [][]string, domainsToMatch map[string]string) map[string]string {
	// printable map to deduplicate
	printableResults := make(map[string]string)
	for _, results := range matches {
		if _, ok := domainsToMatch[results[5]]; ok {
			category := domainsToMatch[results[5]]
			printableResults[results[1]+category+results[5]] = "ip: " + results[1] + ", category: " + category + ", domain: " + results[5]
		}
	}
	return printableResults
}

func printDns(){
	categories := config.Categories.ToCheck
	domainsToMatch := extractDomainsToMatch(categories)
	matches := extractDomainsFromLogs()

	printableResults := deduplicateResults(matches, domainsToMatch)

	for _, info := range printableResults {
		fmt.Printf(info + "\n")
	}
}

func bind9ZonesFormat() string {
	return "zone \"%s\" {type master; file \"/etc/bind/blacklisted.db\";};\n"
}

func updateBlacklistedZones() {
	categories := config.Categories.ToProtect
	domainsToMatch := extractDomainsToMatch(categories)

	writeFileFromStringSlices(domainsToMatch)

	fmt.Println(fmt.Sprintf("%d domains added to blacklist", len(domainsToMatch)))
}

var config Config

func main() {
	config = loadConfig()
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
				Usage:   "update blacklisted.zones domains",
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
		os.Exit(1)
	}

}
