package uno

func (uno *Uno) addPlayer(id string) *Update {
	if uno.state != unoStatePreparation || uno.players.has(id) || unoStartingCardsPerPlayer*(len(uno.players)+1) > 108 {
		return uno.toUpdate()
	}

	uno.players = uno.players.add(newPlayer(id))

	return uno.toUpdate()
}
