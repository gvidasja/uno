package main

type unoPlayer struct {
	name        string
	cards       []*unoCard
	turnsToSkip int
}

func newPlayer(name string) *unoPlayer {
	return &unoPlayer{
		name:        name,
		cards:       []*unoCard{},
		turnsToSkip: 0,
	}
}

func (player *unoPlayer) draw(card *unoCard) {
	player.cards = append(player.cards, card)
}

func (player *unoPlayer) getCard(color string, number string) *unoCard {
	for index, card := range player.cards {
		if card.color == color && card.number == number {
			player.cards = append(player.cards[:index], player.cards[index+1:]...)

			return card
		}
	}

	return nil
}
