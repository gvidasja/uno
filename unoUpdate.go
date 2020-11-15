package main

type UnoUpdate struct {
	Errors []string `json:"error"`
}

func (uno *Uno) ToErrorUpdate(msg string) *UnoUpdate {
	return &UnoUpdate{
		Errors: []string{msg},
	}
}

func (uno *Uno) ToUpdate() *UnoUpdate {
	// TODO implement
	return &UnoUpdate{}
}
