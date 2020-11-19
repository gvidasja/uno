package uno

import "math/rand"

func (uno *Uno) startGame() *Update {
	if len(uno.players) <= 0 {
		return uno.Noop()
	}

	if uno.state != unoStatePreparation {
		return uno.Noop()
	}

	uno.deck = shuffleCards(generateUnoDeck())

	for _, player := range uno.players {
		player.cards = []*unoCard{}
	}

	for i := 0; i < unoStartingCardsPerPlayer; i++ {
		for _, player := range uno.players {
			player.draw(uno.draw())
		}
	}

	uno.winner = nil
	uno.currentPlayerIndex = rand.Intn(len(uno.players) - 1)
	uno.state = unoStatePlaying
	uno.placeCard(uno.draw())

	return uno.toUpdate()
}
