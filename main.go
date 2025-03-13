package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		log.Fatalln("not enough arguments")
	} else if len(args) > 2 {
		log.Fatalln("too many arguments provided")
	}

	region := args[0]
	region = strings.ToLower(region)

	maxConcurrency := 5
	if args[1] != "" {
		num, err := strconv.Atoi(args[1])
		if err != nil {
			log.Fatalf("invalid max concurrency: %s", args[1])
		}
		maxConcurrency = num
	}

	fmt.Printf("starting crawler for crystalline conflict %s\n", region)
	rawBaseURL := "https://na.finalfantasyxiv.com/lodestone/ranking/crystallineconflict/?dcgroup="

	cfg, err := newConfig(rawBaseURL, maxConcurrency)
	if err != nil {
		log.Fatalf("unable to configure: %v", err)
	}

	fmt.Println()
	cfg.scrapePage(region)
	cfg.wg.Wait()

	cfg.sortPlayers()
	cfg.printReport("cc_" + region)
}
