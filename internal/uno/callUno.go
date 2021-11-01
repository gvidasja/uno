package uno

func (uno *Uno) callUno(playerId string) *Update {
	currentPlayer := uno.players[uno.currentPlayerIndex]

	if currentPlayer.id != playerId {
		return uno.toUpdate()
	}

	if len(currentPlayer.cards) != 2 {
		return uno.toUpdate()
	}

	currentPlayer.calledUno = true

	return uno.toUpdate()
}
