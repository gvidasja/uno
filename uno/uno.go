package uno

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

// Uno describes an UNO game
type Uno struct {
	state              string
	players            []*unoPlayer
	deck               []*unoCard
	pile               []*unoCard
	currentPlayerIndex int
	playerOrder        int
	colorOverride      string
}

// Update performs an action on Uno game
func (uno *Uno) Update(action *Action) *Update {
	switch action.Action {
	case unoActionAddPlayer:
		return uno.addPlayer(action.Player)
	case unoActionRemovePlayer:
		return uno.removePlayer(action.Player)
	case unoActionStartGame:
		return uno.startGame()
	case unoActionPlayCard:
		cardColor := action.getDataString(unoActionDataCardColor)
		cardNumber := action.getDataString(unoActionDataCardNumber)
		colorOverride := action.getDataString(unoActionDataColorOverride)
		return uno.playCard(action.Player, cardColor, cardNumber, colorOverride)
	case unoActionDrawCard:
		return uno.drawCard(action.Player)
	case unoActionEndTurn:
	}

	return uno.toErrorUpdate(fmt.Sprint("Unknown action", action.Action))
}

func (uno *Uno) addPlayer(name string) *Update {
	if uno.state != unoStatePreparation && !uno.hasPlayer(name) {
		return uno.toErrorUpdate("Cannot add a new player to started game")
	}

	if uno.hasPlayer(name) {
		// return uno.ToErrorUpdate(fmt.Sprintf("Player %s already exists", name))
		return uno.toUpdate()
	}

	if unoStartingCardsPerPlayer*(len(uno.players)+1) > 108 {
		return uno.toErrorUpdate("Not enough cards for a new player")
	}

	uno.players = append(uno.players, newPlayer(name))

	return uno.toUpdate()
}

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

func (uno *Uno) startGame() *Update {
	if len(uno.players) <= 0 {
		return uno.toErrorUpdate("Cannot start a game with no players")
	}

	if uno.state != unoStatePreparation {
		return uno.toErrorUpdate("Only games in 'PREPARATION' state can be started")
	}

	uno.deck = shuffleCards(generateUnoDeck())

	for _, player := range uno.players {
		player.cards = []*unoCard{}
	}

	for i := 0; i < unoStartingCardsPerPlayer; i++ {
		for _, player := range uno.players {
			player.draw(uno.draw())
		}
	}

	uno.currentPlayerIndex = rand.Intn(len(uno.players) - 1)
	uno.state = unoStatePlaying
	uno.placeCard(uno.draw())

	return uno.toUpdate()
}

func (uno *Uno) playCard(playerName string, cardColor string, cardNumber string, colorOverride string) *Update {
	// TODO quick-move

	currentPlayer := uno.players[uno.currentPlayerIndex]

	if currentPlayer.name != playerName {
		return uno.toErrorUpdate(fmt.Sprintf("It's currently the turn of player %s", currentPlayer.name))
	}

	card := currentPlayer.getCard(cardColor, cardNumber)

	if card == nil {
		return uno.toErrorUpdate(fmt.Sprintf("Player does have card: %s %s", cardColor, cardNumber))
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

	return uno.toUpdate()
}

func (uno *Uno) drawCard(playerName string) *Update {
	currentPlayer := uno.players[uno.currentPlayerIndex]

	if currentPlayer.name != playerName {
		return uno.toErrorUpdate(fmt.Sprintf("It's currently the turn of player %s", currentPlayer.name))
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

func (uno *Uno) placeCard(card *unoCard) {
	uno.pile = append(uno.pile, card)
}

func (uno *Uno) getCurrentTopCard() *unoCard {
	return uno.pile[len(uno.pile)-1]
}

func (uno *Uno) getNextPlayerIndex() int {
	return (uno.currentPlayerIndex + len(uno.players) + uno.playerOrder) % len(uno.players)
}

func (uno *Uno) getPlayer(name string) *unoPlayer {
	for _, player := range uno.players {
		if player.name == name {
			return player
		}
	}

	return nil
}

func (uno *Uno) hasPlayer(name string) bool {
	return uno.getPlayer(name) != nil
}

// NewUno creates a new Uno instance
func NewUno() *Uno {
	return &Uno{
		state:       unoStatePreparation,
		players:     []*unoPlayer{},
		deck:        []*unoCard{},
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
