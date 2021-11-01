package uno

func (uno *Uno) playCard(playerId string, cardColor string, cardNumber string, colorOverride string) *Update {
	// TODO quick-move

	if uno.state != unoStatePlaying {
		return uno.toUpdate()
	}

	currentPlayer := uno.players[uno.currentPlayerIndex]

	if currentPlayer.id != playerId {
		return uno.toUpdate()
	}

	card := currentPlayer.getCard(cardColor, cardNumber)

	if card == nil {
		return uno.toUpdate()
	}

	topCard := uno.getCurrentTopCard()

	if (card.color == unoColorNoColor) || (topCard.color == unoColorNoColor && card.color == uno.colorOverride) || (card.color == topCard.color) || (card.number == topCard.number) {
		uno.colorOverride = ""
		uno.placeCard(card)
		uno.checkIfUnoWasCalled()
		uno.resolveEffect(card, colorOverride)

		uno.resolveCurrentPlayerIndex()

		currentPlayer.canDrawCard = true
	} else {
		currentPlayer.draw(card)
	}

	if len(currentPlayer.cards) == 0 {
		uno.state = unoStateFinished
		uno.winners = append(uno.winners, currentPlayer)
		uno.players, _ = uno.players.remove(currentPlayer.id)
	}

	return uno.toUpdate()
}

func (uno *Uno) resolveEffect(card *unoCard, colorOverride string) {
	switch card.number {
	case unoNumberPlusTwo:
		nextPlayer := uno.players[uno.getNextPlayerIndex()]

		for i := 0; i < 2; i++ {
			nextPlayer.draw(uno.draw())
		}
		nextPlayer.turnsToSkip++
	case unoNumberPlusFour:
		nextPlayer := uno.players[uno.getNextPlayerIndex()]

		for i := 0; i < 4; i++ {
			nextPlayer.draw(uno.draw())
		}
		nextPlayer.turnsToSkip++
		uno.colorOverride = colorOverride
	case unoNumberChangeColor:
		uno.colorOverride = colorOverride
	case unoNumberSkipTurn:
		uno.players[uno.getNextPlayerIndex()].turnsToSkip++
	case unoNumberReverse:
		if len(uno.players) <= 2 {
			uno.players[uno.getNextPlayerIndex()].turnsToSkip++
		} else {
			uno.playerOrder = -uno.playerOrder
		}
	}
}

func (uno *Uno) checkIfUnoWasCalled() {
	currentPlayer := uno.players[uno.currentPlayerIndex]

	if len(currentPlayer.cards) == 1 && !currentPlayer.calledUno {
		currentPlayer.draw(uno.draw())
		currentPlayer.draw(uno.draw())
	}
}
