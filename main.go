package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly/v2"
)

const priceUrl = "https://gg.deals/game/ghost-of-tsushima-directors-cut/"

type GameScrap struct {
	baseUrl string
}

func NewGameScrap() *GameScrap {
	return &GameScrap{baseUrl: "https://gg.deals"}
}

func (gs *GameScrap) Search(gameName string) {
	gameName = gs.sanitize(gameName)
	gs.searchForTheGamePage(gameName)
}

func (gs *GameScrap) searchForTheGamePage(gameName string) {
	//const searchUrl = "https://gg.deals/search/?title="
	searchUrl := fmt.Sprintf("%v/search?title=%v", gs.baseUrl, gameName)
	gameSearchCollector := colly.NewCollector()
	gameSearchCollector.OnHTML(".game-section", func(e *colly.HTMLElement) {
		e.ForEach(".full-link", func(i int, element *colly.HTMLElement) {
			if i == 0 {
				gamePagePath := element.Attr("href")
				gs.searchForGamePrices(gamePagePath)
			}
		})
	})

	gameSearchCollector.OnRequest(func(r *colly.Request) {
		fmt.Printf("visiting %v\n", r.URL)
	})

	gameSearchCollector.Visit(searchUrl)
}
func (gs *GameScrap) searchForGamePrices(gamePagePath string) {
	gamePageCollector := colly.NewCollector()
	gamePageCollector.OnHTML("#official-stores", func(e *colly.HTMLElement) {
		e.ForEach(".game-deals-item", func(_ int, element *colly.HTMLElement) {
			element.ForEach(".game-info-title", func(_ int, element *colly.HTMLElement) {
				fmt.Println(element.Text)
			})
		})
		e.ForEach(".full-link", func(_ int, element *colly.HTMLElement) {
			fmt.Printf("https://gg.deals%v\n", element.Attr("href"))
		})

		e.ForEach(".shop-image-white", func(_ int, element *colly.HTMLElement) {
			fmt.Println(element.Attr("alt"))
		})

		e.ForEach(".game-price-current", func(_ int, element *colly.HTMLElement) {
			fmt.Println(element.Text)
		})
	})

	gamePageCollector.OnRequest(func(r *colly.Request) {
		fmt.Printf("visiting %v\n", r.URL)
	})

	gamePageUrl := fmt.Sprintf("%v%v", gs.baseUrl, gamePagePath)
	gamePageCollector.Visit(gamePageUrl)
}

func (gs *GameScrap) sanitize(gameName string) string {
	gameName = strings.ToLower(gameName)
	gameName = strings.ReplaceAll(gameName, " ", "+")
	return gameName
}

func main() {
	args := os.Args[1:]
	gameName := strings.Join(args, " ")
	//input := "God of war"
	gs := NewGameScrap()
	gs.Search(gameName)
	/*
		c := colly.NewCollector()
		c.OnHTML(".full-link", func(e *colly.HTMLElement) {
			fmt.Println(e.Attr("href"))
		})
		c.Visit(fmt.Sprintf("%v%v", searchUrl, input))
			c.OnHTML("#official-stores", func(e *colly.HTMLElement) {
				e.ForEach(".game-deals-item", func(_ int, element *colly.HTMLElement) {
					element.ForEach(".game-info-title", func(_ int, element *colly.HTMLElement) {
						fmt.Println(element.Text)
					})
				})

				e.ForEach(".full-link", func(_ int, element *colly.HTMLElement) {
					fmt.Printf("https://gg.deals%v\n", element.Attr("href"))
				})

				e.ForEach(".shop-image-white", func(_ int, element *colly.HTMLElement) {
					fmt.Println(element.Attr("alt"))
				})

				e.ForEach(".game-price-current", func(_ int, element *colly.HTMLElement) {
					fmt.Println(element.Text)
				})
			})
	*/

}
