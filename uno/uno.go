package uno

import (
	"log"
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
	winner             *unoPlayer
}

// Execute performs an action on Uno game
func (uno *Uno) Execute(action *Action) *Update {
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

	log.Println("Unkown action:", action)

	return uno.Noop()
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
