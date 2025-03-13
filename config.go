package main

import (
	"fmt"
	"sync"
)

type player struct {
	rank      int
	icon      string
	lodestone string
	name      string
	world     string
	points    string
	wins      string
}

type ranking struct {
	baseURL string
	region  string
	pages   int
	players []player
}

type config struct {
	pages              map[string]int
	rankings           map[string]ranking
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func (cfg *config) printReport(region string) {
	fmt.Println("=============================")
	fmt.Printf("Rankings for %s\n", region)
	fmt.Println("=============================")

	for _, player := range cfg.rankings[region].players {
		fmt.Println(
			player.rank,
			// player.icon,
			// player.lodestone,
			player.name,
			player.world,
			player.points,
			player.wins,
		)
	}
}

func newConfig(rawBaseURL string, maxConcurrency int) (*config, error) {
	rankings := make(map[string]ranking)
	pages := 6

	ccNA := ranking{
		baseURL: rawBaseURL + "Dynamis&page=",
		region:  "cc_na",
		pages:   pages,
		players: []player{},
	}
	ccEU := ranking{
		baseURL: rawBaseURL + "Light&page=",
		region:  "cc_eu",
		pages:   pages,
		players: []player{},
	}
	ccJP := ranking{
		baseURL: rawBaseURL + "Elemental&page=",
		region:  "cc_jp",
		pages:   pages,
		players: []player{},
	}
	ccOCE := ranking{
		baseURL: rawBaseURL + "Materia&page=",
		region:  "cc_oce",
		pages:   pages,
		players: []player{},
	}

	rankings["cc_na"] = ccNA
	rankings["cc_eu"] = ccEU
	rankings["cc_jp"] = ccJP
	rankings["cc_oce"] = ccOCE

	return &config{
		pages:              make(map[string]int),
		rankings:           rankings,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
	}, nil
}
