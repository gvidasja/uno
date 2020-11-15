package main

const (
	unoActionAddPlayer    = "ADD_PLAYER"
	unoActionRemovePlayer = "REMOVE_PLAYER"
	unoActionStartGame    = "START_GAME"
	unoActionPlayCard     = "PLAY_CARD"
	unoActionDrawCard     = "DRAW_CARD"

	unoActionDataPlayerName    = "PLAYER_NAME"
	unoActionDataCardColor     = "CARD_COLOR"
	unoActionDataCardNumber    = "CARD_NUMBER"
	unoActionDataColorOverride = "COLOR_OVERRIDE"
)

type UnoAction struct {
	Action string
	Data   map[string]interface{}
}

func (action *UnoAction) GetDataString(key string) string {
	return action.Data[key].(string)
}
