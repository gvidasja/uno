package uno

func (uno *Uno) removePlayer(id string) *Update {
	if uno.state == unoStatePreparation {
		uno.players, _ = uno.players.remove(id)
	}

	return uno.toUpdate()
}
