package uno

func (uno *Uno) playCard(playerName string, cardColor string, cardNumber string, colorOverride string) *Update {
	// TODO quick-move

	if uno.state != unoStatePlaying {
		return uno.Noop()
	}

	currentPlayer := uno.players[uno.currentPlayerIndex]

	if currentPlayer.name != playerName {
		return uno.Noop()
	}

	card := currentPlayer.getCard(cardColor, cardNumber)

	if card == nil {
		return uno.Noop()
	}

	topCard := uno.getCurrentTopCard()

	if (card.color == unoColorNoColor) || (topCard.color == unoColorNoColor && card.color == uno.colorOverride) || (card.color == topCard.color) || (card.number == topCard.number) {
		uno.colorOverride = ""
		uno.placeCard(card)
		uno.resolveEffect(card, colorOverride)

		uno.resolveCurrentPlayerIndex()

		currentPlayer.canDrawCard = true
	} else {
		currentPlayer.draw(card)
	}

	if len(currentPlayer.cards) == 0 {
		uno.state = unoStateFinished
		uno.winner = currentPlayer
	}

	return uno.toUpdate()
}

func (uno *Uno) resolveEffect(card *unoCard, colorOverride string) {
	switch card.number {
	case unoSpecialCardPlusTwo:
		nextPlayer := uno.players[uno.getNextPlayerIndex()]

		for i := 0; i < 2; i++ {
			nextPlayer.draw(uno.draw())
		}
		nextPlayer.turnsToSkip++
	case unoSpecialCardPlusFour:
		nextPlayer := uno.players[uno.getNextPlayerIndex()]

		for i := 0; i < 4; i++ {
			nextPlayer.draw(uno.draw())
		}
		nextPlayer.turnsToSkip++
		uno.colorOverride = colorOverride
	case unoSpecialCardChangeColor:
		uno.colorOverride = colorOverride
	case unoSpecialCardSkipTurn:
		uno.players[uno.getNextPlayerIndex()].turnsToSkip++
	case unoSpecialCardReverse:
		if len(uno.players) <= 2 {
			uno.players[uno.getNextPlayerIndex()].turnsToSkip++
		} else {
			uno.playerOrder = -uno.playerOrder
		}
	}
}
