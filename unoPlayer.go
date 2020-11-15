package main

type unoPlayer struct {
	name  string
	cards []unoCard
}

func newPlayer(name string) *unoPlayer {
	return &unoPlayer{
		name:  name,
		cards: []unoCard{},
	}
}

func (player *unoPlayer) Draw(uno *Uno) {
	player.cards = append(player.cards, uno.draw())
}
