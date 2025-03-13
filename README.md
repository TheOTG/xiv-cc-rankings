# xiv-cc-rankings

this is a web scraper for the current season of FFXIV Crystalline Conflict

## installation
you will need to install the golang toolchain to run this program
https://webinstall.dev/golang/

## usage
`go run . <region> <concurrency>`
region = na/eu/jp/oce
concurrency = positive number, the higher the faster this will fetch the data

e.g. `go run . na 6`

### wip
- getting the links from the base url instead of having to hard code the datacenters