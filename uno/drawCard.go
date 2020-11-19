package uno

func (uno *Uno) drawCard(playerName string) *Update {
	if uno.state != unoStatePlaying {
		return uno.Noop()
	}

	currentPlayer := uno.players[uno.currentPlayerIndex]

	if currentPlayer.name != playerName {
		return uno.Noop()
	}

	if currentPlayer.canDrawCard {
		currentPlayer.draw(uno.draw())
		currentPlayer.canDrawCard = false
	} else {
		currentPlayer.canDrawCard = true
		uno.resolveCurrentPlayerIndex()
	}

	return uno.toUpdate()
}
