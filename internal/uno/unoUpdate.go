package uno

import (
	"fmt"
	"sort"
)

// UpdatePlayer is a player DTO
type UpdatePlayer struct {
	Id       string `json:"id"`
	Turn     bool   `json:"turn"`
	HandSize int    `json:"handSize"`
	Winner   bool   `json:"winner"`
}

// UpdateCard is a card DTO
type UpdateCard struct {
	Color  string `json:"color"`
	Number string `json:"number"`
}

// UpdateForPlayer is a player-specific wrapper for update DTO
type UpdateForPlayer struct {
	GlobalState *Update       `json:"globalState"`
	Me          *UpdatePlayer `json:"me"`
	Hand        []*UpdateCard `json:"hand"`
}

// Update is a game update DTO
type Update struct {
	State         string          `json:"state"`
	Players       []*UpdatePlayer `json:"players"`
	PileSize      int             `json:"pileSize"`
	DeckSize      int             `json:"deckSize"`
	TopCard       *UpdateCard     `json:"topCard"`
	ColorOverride string          `json:"colorOverride"`
	Errors        []string        `json:"errors"`
	Winner        *UpdatePlayer   `json:"winner"`
	uno           *Uno
}

func (uno *Uno) toUpdate() *Update {
	players := []*UpdatePlayer{}

	allPlayers := append(uno.players, uno.winners...)
	sort.Slice(allPlayers, func(i int, j int) bool {
		return allPlayers[i].index < allPlayers[j].index
	})

	for index, player := range allPlayers {
		players = append(players, &UpdatePlayer{
			Id:       player.id,
			HandSize: len(player.cards),
			Turn:     index == uno.currentPlayerIndex,
			Winner:   uno.winners.has(player.id),
		})
	}

	var topCard *UpdateCard

	if uno.state == unoStatePreparation {
		topCard = nil
	} else {
		topCardModel := uno.getCurrentTopCard()
		topCard = &UpdateCard{
			Color:  topCardModel.color,
			Number: topCardModel.number,
		}
	}

	return &Update{
		State:         uno.state,
		Players:       players,
		ColorOverride: uno.colorOverride,
		TopCard:       topCard,
		DeckSize:      len(uno.deck),
		PileSize:      len(uno.pile),
		uno:           uno,
	}
}

// ForPlayer wraps an update DTO with player-specific info
func (update *Update) ForPlayer(id string) *UpdateForPlayer {
	hand := []*UpdateCard{}
	player := update.uno.players.get(id)

	if player == nil || player.cards == nil && !update.uno.winners.has(id) {
		panic(fmt.Sprint("Player does not exist:", id))
	}

	if !update.uno.winners.has(id) {
		for _, card := range player.cards {
			hand = append(hand, &UpdateCard{
				Color:  card.color,
				Number: card.number,
			})
		}
	}

	sort.Slice(hand, func(i int, j int) bool {
		return hand[i].Color < hand[j].Color ||
			(hand[i].Color == hand[j].Color &&
				hand[i].Number < hand[j].Number)
	})

	var me *UpdatePlayer

	for _, player := range update.Players {
		if player.Id == id {
			me = player
			break
		}
	}

	return &UpdateForPlayer{
		GlobalState: update,
		Hand:        hand,
		Me:          me,
	}
}
