package uno

type Players []*unoPlayer

func (players Players) get(id string) *unoPlayer {
	for _, p := range players {
		if p.id == id {
			return p
		}
	}

	return nil
}

func (players Players) remove(id string) (Players, *unoPlayer) {
	newPlayers := Players{}
	var removedPlayer *unoPlayer

	for _, p := range players {
		if p.id == id {
			removedPlayer = p
		} else {
			newPlayers = append(newPlayers, p)
		}
	}

	return newPlayers, removedPlayer
}

func (players Players) add(player *unoPlayer) Players {
	return append(players, player)
}

func (players Players) has(id string) bool {
	return players.get(id) != nil
}
