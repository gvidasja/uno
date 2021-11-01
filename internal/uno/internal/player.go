package player

type Player struct {
	id           string
	name         string
	hasToSkip    bool
	canDraw      bool
	hasToDraw    int
	hasToCallUno bool
	playing      bool
}

func New(id string, name string) *Player {
	return &Player{
		id:           id,
		name:         name,
		hasToSkip:    false,
		canDraw:      false,
		hasToDraw:    0,
		hasToCallUno: false,
		playing:      true,
	}
}
