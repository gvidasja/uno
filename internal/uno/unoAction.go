package uno

type UnoAction string
type UnoActionData string

const (
	unoActionAddPlayer    = UnoAction("ADD_PLAYER")
	unoActionRemovePlayer = UnoAction("REMOVE_PLAYER")
	unoActionStartGame    = UnoAction("START_GAME")
	unoActionPerformMove  = UnoAction("PERFORM_MOVE")

	unoActionPlayCard      = UnoAction("PLAY_CARD")
	unoActionDrawCard      = UnoAction("DRAW_CARD")
	unoActionChangeColor   = UnoAction("CHANGE_COLOR")
	unoActionDrawTwoCards  = UnoAction("DRAW_TWO_CARDS")
	unoActionDrawFourCards = UnoAction("DRAW_FOUNR_CARDS")
	unoActionEndTurn       = UnoAction("END_TURN")
	unoActionCallUno       = UnoAction("CALL_UNO")

	unoActionDataCardColor     = UnoActionData("CARD_COLOR")
	unoActionDataCardNumber    = UnoActionData("CARD_NUMBER")
	unoActionDataColorOverride = UnoActionData("COLOR_OVERRIDE")
)

// Action describes an action that can be performed in a game
type Action struct {
	Action UnoAction
	Player string
	Data   map[UnoActionData]interface{}
}

func (action *Action) getDataString(key UnoActionData) string {
	if action.Data == nil || action.Data[key] == nil {
		return ""
	}

	value, ok := action.Data[key].(string)

	if !ok {
		return ""
	}

	return value
}
