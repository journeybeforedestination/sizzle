package pan

type Pan struct {
	Sizzles []string
}

func (p *Pan) TurnItOn() {
	p.Sizzles = []string{"123", "456"}
}
