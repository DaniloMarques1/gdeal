package main

import (
	"fmt"
	"sync"
)

type GamePrice struct {
	mut   sync.Mutex
	games map[int]Game
}

func NewGamePrice() *GamePrice {
	games := make(map[int]Game)
	return &GamePrice{games: games}
}

func (gp *GamePrice) AddGameTitle(idx int, title string) {
	gp.mut.Lock()
	defer gp.mut.Unlock()

	game, ok := gp.games[idx]
	if !ok {
		gp.games[idx] = Game{title: title}
	} else {
		game.title = title
		gp.games[idx] = game
	}
}

func (gp *GamePrice) AddGamePrice(idx int, price string) {
	gp.mut.Lock()
	defer gp.mut.Unlock()

	game, ok := gp.games[idx]
	if !ok {
		gp.games[idx] = Game{price: price}
	} else {
		game.price = price
		gp.games[idx] = game
	}
}

func (gp *GamePrice) AddGameShopName(idx int, shopName string) {
	gp.mut.Lock()
	defer gp.mut.Unlock()

	game, ok := gp.games[idx]
	if !ok {
		gp.games[idx] = Game{shopName: shopName}
	} else {
		game.shopName = shopName
		gp.games[idx] = game
	}
}

func (gp *GamePrice) AddGameShopUrl(idx int, shopUrl string) {
	gp.mut.Lock()
	defer gp.mut.Unlock()

	game, ok := gp.games[idx]
	if !ok {
		gp.games[idx] = Game{shopName: shopUrl}
	} else {
		game.shopUrl = shopUrl
		gp.games[idx] = game
	}
}

func (gp *GamePrice) Print() {
	fmt.Println("===============================================")
	for i := 0; i < len(gp.games); i++ {
		game := gp.games[i]
		fmt.Printf("Game: %v\n", game.title)
		fmt.Printf("Store: %v\n", game.shopName)
		fmt.Printf("Price: %v\n", game.price)
		fmt.Printf("Shop: %v\n", game.shopUrl)
	}
	fmt.Println("===============================================")
}

type Game struct {
	title    string
	price    string
	shopUrl  string
	shopName string
}
