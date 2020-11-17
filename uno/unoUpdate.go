package uno

import (
	"fmt"
	"sort"
)

// UpdatePlayer is a player DTO
type UpdatePlayer struct {
	Name     string `json:"name"`
	Turn     bool   `json:"turn"`
	HandSize int    `json:"handSize"`
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
	uno           *Uno
}

func (uno *Uno) toErrorUpdate(msg string) *Update {
	return &Update{
		Errors: []string{msg},
		uno:    uno,
	}
}

func (uno *Uno) toUpdate() *Update {
	players := []*UpdatePlayer{}

	for index, player := range uno.players {
		players = append(players, &UpdatePlayer{
			Name:     player.name,
			HandSize: len(player.cards),
			Turn:     index == uno.currentPlayerIndex,
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
func (update *Update) ForPlayer(name string) *UpdateForPlayer {
	hand := []*UpdateCard{}
	player := update.uno.getPlayer(name)

	if player == nil || player.cards == nil {
		panic(fmt.Sprint("Player does not exist:", name))
	}

	for _, card := range player.cards {
		hand = append(hand, &UpdateCard{
			Color:  card.color,
			Number: card.number,
		})
	}

	sort.Slice(hand, func(i int, j int) bool {
		return hand[i].Color < hand[j].Color ||
			(hand[i].Color == hand[j].Color &&
				hand[i].Number < hand[j].Number)
	})

	var me *UpdatePlayer

	for _, player := range update.Players {
		if player.Name == name {
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
