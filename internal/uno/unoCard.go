package uno

import (
	"math/rand"
	"time"
)

type unoColor string
type unoNumber string

const (
	unoColorRed     = unoColor("RED")
	unoColorGreen   = unoColor("GREEN")
	unoColorBlue    = unoColor("BLUE")
	unoColorYellow  = unoColor("YELLOW")
	unoColorNoColor = unoColor("NO_COLOR")

	unoNumber0           = unoNumber("0")
	unoNumber1           = unoNumber("1")
	unoNumber2           = unoNumber("2")
	unoNumber3           = unoNumber("3")
	unoNumber4           = unoNumber("4")
	unoNumber5           = unoNumber("5")
	unoNumber6           = unoNumber("6")
	unoNumber7           = unoNumber("7")
	unoNumber8           = unoNumber("8")
	unoNumber9           = unoNumber("9")
	unoNumberReverse     = unoNumber("REVERSE")
	unoNumberPlusTwo     = unoNumber("PLUS_TWO")
	unoNumberSkipTurn    = unoNumber("SKIP_TURN")
	unoNumberPlusFour    = unoNumber("PLUS_FOUR")
	unoNumberChangeColor = unoNumber("CHANGE_COLOR")
)

type unoCard struct {
	color         unoColor
	number        unoNumber
	overrideColor unoColor
}

func (c *unoCard) canBeSucceededBy(card *unoCard) bool {
	return (c.color == unoColorNoColor && c.color == card.color && c.number == card.number) ||
		(c.color == card.color || c.number == card.number) ||
		(c.overrideColor)
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
		cardNumbers := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", unoNumberPlusTwo, unoNumberReverse, unoNumberSkipTurn}

		for _, cardNumber := range cardNumbers {
			card := &unoCard{color, cardNumber}
			cards = append(cards, repeat(card, 2)...)
		}
	}

	plusFourCard := &unoCard{unoColorNoColor, unoNumberPlusFour}
	changeColorCard := &unoCard{unoColorNoColor, unoNumberChangeColor}

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
