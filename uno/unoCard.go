package uno

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

func shuffleCards(cards []*unoCard) []*unoCard {
	random := rand.New(rand.NewSource(time.Now().Unix()))
	shuffledCards := cards[:]

	random.Shuffle(len(cards), func(i int, j int) {
		shuffledCards[i], shuffledCards[j] = shuffledCards[j], shuffledCards[i]
	})

	return shuffledCards
}

func generateUnoDeck() []*unoCard {
	cards := []*unoCard{}
	colors := []string{unoColorRed, unoColorGreen, unoColorBlue, unoColorYellow}

	for _, color := range colors {
		cards = append(cards, &unoCard{color, "0"})
		cardNumbers := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", unoSpecialCardPlusTwo, unoSpecialCardReverse, unoSpecialCardSkipTurn}

		for _, cardNumber := range cardNumbers {
			card := &unoCard{color, cardNumber}
			cards = append(cards, repeat(card, 2)...)
		}
	}

	plusFourCard := &unoCard{unoColorNoColor, unoSpecialCardPlusFour}
	changeColorCard := &unoCard{unoColorNoColor, unoSpecialCardChangeColor}

	cards = append(cards, repeat(plusFourCard, 4)...)
	cards = append(cards, repeat(changeColorCard, 4)...)

	return cards
}

func repeat(card *unoCard, times int) []*unoCard {
	cards := make([]*unoCard, times)

	for i := 0; i < times; i++ {
		cards[i] = card
	}

	return cards
}
