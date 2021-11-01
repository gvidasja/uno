package uno

func (uno *Uno) drawCard(playerId string) *Update {
	if uno.state != unoStatePlaying {
		return uno.toUpdate()
	}

	currentPlayer := uno.players[uno.currentPlayerIndex]

	if currentPlayer.id != playerId || currentPlayer.out {
		return uno.toUpdate()
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
