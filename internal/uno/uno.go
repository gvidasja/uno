package uno

import (
	"log"
)

type unoState string

const (
	unoStatePreparation = unoState("PREPARATION")
	unoStatePlaying     = unoState("PLAYING")
	unoStateFinished    = unoState("FINISHED")

	unoStartingCardsPerPlayer = 7
)

type Events interface {
	Send(playerID string, u Uno)
}

// Uno describes an UNO game
type Uno struct {
	state              unoState
	players            Players
	deck               []*unoCard
	pile               []*unoCard
	currentPlayerIndex int
	playerOrder        int

	events Events
}

func (u *Uno) Add(playerID string, playerName string) {
	if u.state == unoStatePreparation {
		u.players.add(newPlayer(playerID))
	}
}

// Execute performs an action on Uno game
func (u *Uno) Execute(action *Action) *Update {
	switch action.Action {
	case unoActionAddPlayer:
		return u.addPlayer(action.Player)
	case unoActionRemovePlayer:
		return u.removePlayer(action.Player)
	case unoActionStartGame:
		return u.startGame()
	case unoActionPlayCard:
		cardColor := action.getDataString(unoActionDataCardColor)
		cardNumber := action.getDataString(unoActionDataCardNumber)
		colorOverride := action.getDataString(unoActionDataColorOverride)
		return u.playCard(action.Player, cardColor, cardNumber, colorOverride)
	case unoActionDrawCard:
		return u.drawCard(action.Player)
	case unoActionCallUno:
		return u.callUno(action.Player)
	case unoActionEndTurn:
	}

	log.Println("Unkown action:", action)

	return u.toUpdate()
}

func (u *Uno) resolveCurrentPlayerIndex() {
	for {
		u.currentPlayerIndex = u.getNextPlayerIndex()

		if u.players[u.currentPlayerIndex].turnsToSkip <= 0 {
			break
		} else {
			u.players[u.currentPlayerIndex].turnsToSkip--
		}
	}
}

func (u *Uno) placeCard(card *unoCard) {
	u.pile = append(u.pile, card)
}

func (u *Uno) getCurrentTopCard() *unoCard {
	return u.pile[len(u.pile)-1]
}

func (u *Uno) getNextPlayerIndex() int {
	return (u.currentPlayerIndex + len(u.players) + u.playerOrder) % len(u.players)
}

// NewUno creates a new Uno instance
func NewUno() *Uno {
	return &Uno{
		state:       unoStatePreparation,
		players:     []*unoPlayer{},
		deck:        []*unoCard{},
		pile:        []*unoCard{},
		playerOrder: 1,
		winners:     []*unoPlayer{},
	}
}

func (u *Uno) draw() *unoCard {
	if len(u.deck) <= 0 {
		u.deck = shuffleCards(u.pile[:len(u.pile)-1])
		u.pile = u.pile[len(u.pile)-1:]
	}

	drawnCard := u.deck[len(u.deck)-1]
	u.deck = u.deck[:len(u.deck)-1]

	return drawnCard
}
