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
	deck               []unoCard
	pile               []unoCard
	currentPlayerIndex int
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
		return uno.playCard(playerName, cardColor, cardNumber)
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

	if len(uno.deck) < unoStartingCardsPerPlayer {
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

	return uno.ToUpdate()
}

func (uno *Uno) startGame() *UnoUpdate {
	if len(uno.players) <= 0 {
		return uno.ToErrorUpdate("Cannot start a game with no players")
	}

	for i := 0; i < unoStartingCardsPerPlayer; i++ {
		for _, player := range uno.players {
			player.Draw(uno)
		}
	}

	uno.currentPlayerIndex = rand.Intn(len(uno.players) - 1)
	uno.state = unoStatePlaying

	return uno.ToUpdate()
}

func (uno *Uno) playCard(playerName string, cardColor string, cardNumber string) *UnoUpdate {
	// TODO quick-move

	currentPlayer := uno.players[uno.currentPlayerIndex]

	if currentPlayer.name != playerName {
		return uno.ToErrorUpdate(fmt.Sprintf("It's currently the turn of player %s", currentPlayer.name))
	}

	// TODO implement card play

	return uno.ToUpdate()
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
		state:   unoStatePreparation,
		players: []*unoPlayer{},
		deck:    shuffleCards(generateUnoDeck()),
		pile:    []unoCard{},
	}
}

func (uno *Uno) draw() unoCard {
	if len(uno.deck) <= 0 {
		uno.deck = shuffleCards(uno.pile[:len(uno.pile)-1])
		uno.pile = uno.pile[len(uno.pile)-1:]
	}

	drawnCard := uno.deck[len(uno.deck)-1]
	uno.deck = uno.deck[:len(uno.deck)-1]

	return drawnCard
}
