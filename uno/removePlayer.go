package uno

func (uno *Uno) removePlayer(name string) *Update {
	newPlayers := []*unoPlayer{}

	for _, player := range uno.players {
		if player.name != name {
			newPlayers = append(newPlayers, player)
		}
	}

	uno.players = newPlayers
	uno.currentPlayerIndex = uno.currentPlayerIndex + len(uno.players)

	return uno.toUpdate()
}
