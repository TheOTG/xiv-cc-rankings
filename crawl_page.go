package main

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
)

func (cfg *config) scrapePage(reg string) {
	for region, ranking := range cfg.rankings {
		if region != "cc_"+reg {
			continue
		}
		fmt.Printf("crawling %s\n", region)
		for i := range ranking.pages {
			cfg.wg.Add(1)
			go func() {
				cfg.concurrencyControl <- struct{}{}
				defer func() {
					<-cfg.concurrencyControl
					cfg.wg.Done()
				}()
				html, err := getHTML(ranking.baseURL + strconv.Itoa(i+1))
				if err != nil {
					fmt.Printf("unable to get html: %v", err)
					return
				}

				players, err := getPlayersFromHTML(html)
				if err != nil {
					fmt.Printf("unable to get urls: %v", err)
					return
				}

				cfg.addPlayers(region, players)
			}()
		}
	}

}

func sortPlayerFunc(a, b player) int {
	return cmp.Compare(a.rank, b.rank)
}

func (cfg *config) addPlayers(region string, players []player) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	ranking, ok := cfg.rankings[region]
	if ok {
		ranking.players = append(ranking.players, players...)
		cfg.rankings[region] = ranking
	}
}

func (cfg *config) sortPlayers() {
	for _, ranking := range cfg.rankings {
		slices.SortFunc(ranking.players, sortPlayerFunc)
	}
}
