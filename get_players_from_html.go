package main

import (
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func getPlayersFromHTML(htmlBody string) ([]player, error) {
	htmlReader := strings.NewReader(htmlBody)
	doc, err := html.Parse(htmlReader)
	if err != nil {
		return []player{}, fmt.Errorf("couldn't parse HTML: %v", err)
	}

	rankings := []player{}

	for n := range doc.Descendants() {
		if n.Type == html.ElementNode && n.DataAtom == atom.Div {
			for _, a := range n.Attr {
				if a.Key == "class" && strings.Contains(a.Val, "ranking_set") {
					p := player{}

					p.lodestone = n.Attr[1].Val

					for n2 := range n.Descendants() {
						if n2.Type == html.ElementNode && n2.DataAtom == atom.Div {
							for _, a2 := range n2.Attr {
								if a2.Key == "class" && a2.Val == "order" {
									p.rank, err = strconv.Atoi(strings.Trim(n2.FirstChild.Data, "\n\t"))
									if err != nil {
										return []player{}, fmt.Errorf("unable to convert rank: %v", err)
									}
									break
								}
								if a2.Key == "class" && a2.Val == "face-wrapper" {
									for n3 := range n2.Descendants() {
										if n3.Type == html.ElementNode && n3.DataAtom == atom.Img {
											p.icon = n3.Attr[0].Val
											break
										}
									}
									break
								}
								if a2.Key == "class" && a2.Val == "cc-ranking__result__name" {
									for n3 := range n2.Descendants() {
										if n3.Type == html.ElementNode && n3.DataAtom == atom.H3 {
											p.name = n3.FirstChild.Data
										}
										if n3.Type == html.ElementNode && n3.DataAtom == atom.Span {
											p.world = n3.LastChild.Data
										}
									}
								}
								if a2.Key == "class" && a2.Val == "points" {
									for n3 := range n2.Descendants() {
										if n3.Type == html.ElementNode && n3.DataAtom == atom.P {
											p.points = n3.FirstChild.Data
										}
										if n3.Type == html.ElementNode && n3.DataAtom == atom.Span {
											p.points += n3.FirstChild.Data
										}
									}
								}
								if a2.Key == "class" && a2.Val == "wins" {
									for n3 := range n2.Descendants() {
										if n3.Type == html.ElementNode && n3.DataAtom == atom.P {
											p.wins = n3.FirstChild.Data
										}
										if n3.Type == html.ElementNode && n3.DataAtom == atom.Span {
											p.wins += n3.FirstChild.Data
										}
									}
								}
							}
						}
					}

					if p.name != "" {
						rankings = append(rankings, p)
						break
					}
					break
				}
			}
		}
	}
	return rankings, nil
}
