package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

type GameScrap struct {
	baseUrl   string
	gameName  string
	gamePrice *GamePrice
}

func NewGameScrap(gameName string) *GameScrap {
	gs := GameScrap{baseUrl: "https://gg.deals"}
	gs.gameName = gs.sanitize(gameName)
	gs.gamePrice = NewGamePrice()
	return &gs
}

func (gs *GameScrap) Search() {
	gs.searchForTheGamePage()
}

func (gs *GameScrap) searchForTheGamePage() {
	searchUrl := fmt.Sprintf("%v/search?title=%v", gs.baseUrl, gs.gameName)
	gameSearchCollector := colly.NewCollector()
	gameSearchCollector.OnHTML(".game-section", func(e *colly.HTMLElement) {
		e.ForEach(".full-link", func(i int, element *colly.HTMLElement) {
			if i == 0 {
				gamePagePath := element.Attr("href")
				gs.searchForGamePrices(gamePagePath)
			}
		})
	})

	//gameSearchCollector.OnRequest(func(r *colly.Request) {
	//	fmt.Printf("visiting %v\n", r.URL)
	//})

	gameSearchCollector.Visit(searchUrl)
}
func (gs *GameScrap) searchForGamePrices(gamePagePath string) {
	gamePageCollector := colly.NewCollector()
	gamePageCollector.OnHTML("#official-stores", func(e *colly.HTMLElement) {
		e.ForEach(".game-info-title", func(idx int, element *colly.HTMLElement) {
			//fmt.Printf("%v - %v\n", idx, element.Text)
			gs.gamePrice.AddGameTitle(idx, element.Text)
		})

		e.ForEach(".full-link", func(idx int, element *colly.HTMLElement) {
			shopUrl := fmt.Sprintf("https://gg.deals%v\n", element.Attr("href"))
			gs.gamePrice.AddGameShopUrl(idx, shopUrl)
		})

		e.ForEach(".shop-image-white", func(idx int, element *colly.HTMLElement) {
			shopName := element.Attr("alt")
			gs.gamePrice.AddGameShopName(idx, shopName)
		})

		e.ForEach(".game-price-current", func(idx int, element *colly.HTMLElement) {
			gs.gamePrice.AddGamePrice(idx, element.Text)
		})
	})

	//gamePageCollector.OnRequest(func(r *colly.Request) {
	//	fmt.Printf("visiting %v\n", r.URL)
	//})

	gamePageCollector.OnScraped(func(r *colly.Response) {
		gs.gamePrice.Print()
	})

	gamePageUrl := fmt.Sprintf("%v%v", gs.baseUrl, gamePagePath)
	gamePageCollector.Visit(gamePageUrl)
}

func (gs *GameScrap) sanitize(gameName string) string {
	gameName = strings.ToLower(gameName)
	gameName = strings.ReplaceAll(gameName, " ", "+")
	return gameName
}
