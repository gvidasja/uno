package uno

type unoPlayer struct {
	id    string
	index int
	cards []*unoCard

	turnsToSkip int
	canDrawCard bool
	calledUno   bool
	out         bool
}

func newPlayer(id string) *unoPlayer {
	return &unoPlayer{
		id:          id,
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
