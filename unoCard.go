package main

import (
	"math/rand"
	"time"
)

const (
	unoColorRed     = "RED"
	unoColorGreen   = "GREEN"
	unoColorBlue    = "BLUE"
	unoColorYellow  = "YELLOW"
	unoColorNoColor = "NO_COLOR"

	unoSpecialCardReverse     = "REVERSE"
	unoSpecialCardPlusTwo     = "PLUS_TWO"
	unoSpecialCardSkipTurn    = "SKIP_TURN"
	unoSpecialCardPlusFour    = "PLUS_FOUR"
	unoSpecialCardChangeColor = "CHANGE_COLOR"
)

type unoCard struct {
	color  string
	number string
}

func shuffleCards(cards []unoCard) []unoCard {
	random := rand.New(rand.NewSource(time.Now().Unix()))
	shuffledCards := make([]unoCard, len(cards))

	for len(cards) > 0 {
		randomCardIndex := random.Intn(len(cards))
		shuffledCards = append(shuffledCards, cards[randomCardIndex])
		cards = append(cards[:randomCardIndex], cards[randomCardIndex+1:]...)
	}

	return shuffledCards
}

func generateUnoDeck() []unoCard {
	cards := make([]unoCard, 108)
	colors := []string{unoColorRed, unoColorGreen, unoColorBlue, unoColorYellow}

	for _, color := range colors {
		cards = append(cards, unoCard{color, "0"})
		cardNumbers := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", unoSpecialCardPlusTwo, unoSpecialCardReverse, unoSpecialCardSkipTurn}

		for _, cardNumber := range cardNumbers {
			card := unoCard{color, cardNumber}
			cards = append(cards, card)
		}
	}

	plusFourCard := unoCard{unoColorNoColor, unoSpecialCardPlusFour}
	changeColorCard := unoCard{unoColorNoColor, unoSpecialCardChangeColor}

	cards = append(cards, repeat(plusFourCard, 4)...)
	cards = append(cards, repeat(changeColorCard, 4)...)

	return cards
}

func repeat(card unoCard, times int) []unoCard {
	cards := make([]unoCard, times)

	for i := 0; i < times; i++ {
		cards = append(cards, card)
	}

	return cards
}
