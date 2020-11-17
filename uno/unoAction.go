package uno

const (
	unoActionAddPlayer    = "ADD_PLAYER"
	unoActionRemovePlayer = "REMOVE_PLAYER"
	unoActionStartGame    = "START_GAME"
	unoActionPlayCard     = "PLAY_CARD"
	unoActionDrawCard     = "DRAW_CARD"
	unoActionEndTurn      = "END_TURN"

	unoActionDataCardColor     = "CARD_COLOR"
	unoActionDataCardNumber    = "CARD_NUMBER"
	unoActionDataColorOverride = "COLOR_OVERRIDE"
)

// Action describes an action that can be performed in a game
type Action struct {
	Action string
	Player string
	Data   map[string]interface{}
}

func (action *Action) getDataString(key string) string {
	if action.Data == nil || action.Data[key] == nil {
		return ""
	}

	value, ok := action.Data[key].(string)

	if !ok {
		return ""
	}

	return value
}
