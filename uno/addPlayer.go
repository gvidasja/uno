package uno

func (uno *Uno) addPlayer(name string) *Update {
	if uno.state != unoStatePreparation && !uno.hasPlayer(name) {
		return uno.Noop()
	}

	if uno.hasPlayer(name) {
		// return uno.ToErrorUpdate(fmt.Sprintf("Player %s already exists", name))
		return uno.toUpdate()
	}

	if unoStartingCardsPerPlayer*(len(uno.players)+1) > 108 {
		return uno.Noop()
	}

	uno.players = append(uno.players, newPlayer(name))

	return uno.toUpdate()
}
