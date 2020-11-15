package main

import (
	"fmt"
	"math/rand"
)

const (
	unoStatePreparation = "PREPARATION"
	unoStatePlaying     = "PLAYING"
	unoStateFinished    = "FINISHED"

	unoStartingCardsPerPlayer = 7
)

type Uno struct {
	state              string
	players            []*unoPlayer
	deck               []*unoCard
	pile               []*unoCard
	currentPlayerIndex int
	playerOrder        int
	colorOverride      string
}

func (uno *Uno) Update(action *UnoAction) *UnoUpdate {
	switch action.Action {
	case unoActionAddPlayer:
		playerName := action.GetDataString(unoActionDataPlayerName)
		return uno.addPlayer(playerName)
	case unoActionRemovePlayer:
		playerName := action.GetDataString(unoActionDataPlayerName)
		return uno.removePlayer(playerName)
	case unoActionStartGame:
		return uno.startGame()
	case unoActionPlayCard:
		playerName := action.GetDataString(unoActionDataPlayerName)
		cardColor := action.GetDataString(unoActionDataCardColor)
		cardNumber := action.GetDataString(unoActionDataCardNumber)
		colorOverride := action.GetDataString(unoActionDataColorOverride)
		return uno.playCard(playerName, cardColor, cardNumber, colorOverride)
	}

	return &UnoUpdate{}
}

func (uno *Uno) addPlayer(name string) *UnoUpdate {
	if uno.state != unoStatePreparation {
		return uno.ToErrorUpdate("Cannot add player to started game")
	}

	if uno.hasPlayer(name) {
		return uno.ToErrorUpdate(fmt.Sprintf("Player %s already exists", name))
	}

	if len(uno.deck) < unoStartingCardsPerPlayer*2 {
		return uno.ToErrorUpdate("Not enough cards for a new player")
	}

	uno.players = append(uno.players, newPlayer(name))

	return uno.ToUpdate()
}

func (uno *Uno) removePlayer(name string) *UnoUpdate {
	newPlayers := make([]*unoPlayer, len(uno.players))

	for _, player := range uno.players {
		if player.name != name {
			newPlayers = append(newPlayers, player)
		}
	}

	uno.players = newPlayers
	uno.currentPlayerIndex = uno.currentPlayerIndex + len(uno.players)

	return uno.ToUpdate()
}

func (uno *Uno) startGame() *UnoUpdate {
	if len(uno.players) <= 0 {
		return uno.ToErrorUpdate("Cannot start a game with no players")
	}

	for i := 0; i < unoStartingCardsPerPlayer; i++ {
		for _, player := range uno.players {
			player.draw(uno.draw())
		}
	}

	uno.currentPlayerIndex = rand.Intn(len(uno.players) - 1)
	uno.state = unoStatePlaying
	uno.placeCard(uno.draw())

	return uno.ToUpdate()
}

func (uno *Uno) playCard(playerName string, cardColor string, cardNumber string, colorOverride string) *UnoUpdate {
	// TODO quick-move

	currentPlayer := uno.players[uno.currentPlayerIndex]

	if currentPlayer.name != playerName {
		return uno.ToErrorUpdate(fmt.Sprintf("It's currently the turn of player %s", currentPlayer.name))
	}

	card := currentPlayer.getCard(cardColor, cardNumber)

	if card == nil {
		return uno.ToErrorUpdate(fmt.Sprintf("Player does have card: %s %s", cardColor, cardNumber))
	}

	topCard := uno.getCurrentTopCard()

	if card.color == unoColorNoColor || (topCard.color == unoColorNoColor && topCard.color == uno.colorOverride) || card.color == topCard.color || card.number == card.number {
		uno.colorOverride = ""
		uno.placeCard(card)
		uno.resolveEffect(card, colorOverride)

		uno.resolveCurrentPlayerIndex()
	} else {
		currentPlayer.draw(card)
	}

	return uno.ToUpdate()
}

func (uno *Uno) resolveCurrentPlayerIndex() {
	for {
		uno.currentPlayerIndex = uno.getNextPlayerIndex()

		if uno.players[uno.currentPlayerIndex].turnsToSkip <= 0 {
			break
		} else {
			uno.players[uno.currentPlayerIndex].turnsToSkip--
		}
	}
}

func (uno *Uno) resolveEffect(card *unoCard, colorOverride string) {
	switch card.number {
	case unoSpecialCardPlusTwo:
		nextPlayer := uno.players[uno.getNextPlayerIndex()]

		for i := 0; i < 2; i++ {
			nextPlayer.draw(uno.draw())
			nextPlayer.turnsToSkip++
		}
	case unoSpecialCardPlusFour:
		nextPlayer := uno.players[uno.getNextPlayerIndex()]

		for i := 0; i < 4; i++ {
			nextPlayer.draw(uno.draw())
			nextPlayer.turnsToSkip++
		}

		uno.colorOverride = colorOverride
	case unoSpecialCardChangeColor:
		uno.colorOverride = colorOverride
	case unoSpecialCardSkipTurn:
		uno.players[uno.getNextPlayerIndex()].turnsToSkip++
	case unoSpecialCardReverse:
		uno.playerOrder = -uno.playerOrder
	}
}

func (uno *Uno) placeCard(card *unoCard) {
	uno.pile = append(uno.pile, card)
}

func (uno *Uno) getCurrentTopCard() *unoCard {
	return uno.pile[len(uno.pile)-1]
}

func (uno *Uno) getNextPlayerIndex() int {
	return (uno.currentPlayerIndex + len(uno.players) + uno.playerOrder) % len(uno.players)
}

func (uno *Uno) hasPlayer(name string) bool {
	for _, player := range uno.players {
		if player.name == name {
			return true
		}
	}

	return false
}

func NewUno() *Uno {
	return &Uno{
		state:       unoStatePreparation,
		players:     []*unoPlayer{},
		deck:        shuffleCards(generateUnoDeck()),
		pile:        []*unoCard{},
		playerOrder: 1,
	}
}

func (uno *Uno) draw() *unoCard {
	if len(uno.deck) <= 0 {
		uno.deck = shuffleCards(uno.pile[:len(uno.pile)-1])
		uno.pile = uno.pile[len(uno.pile)-1:]
	}

	drawnCard := uno.deck[len(uno.deck)-1]
	uno.deck = uno.deck[:len(uno.deck)-1]

	return drawnCard
}
