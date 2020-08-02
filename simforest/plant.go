package simforest

type Plant struct {
	Position
	isAlive bool
}

func (p *Plant) Act(population []Entity) []Entity {
	return []Entity{}
}

func (p *Plant) Gender() Gender {
	return Female
}

func (p *Plant) IsAdult() bool {
	return false
}

func (p *Plant) IsAlive() bool {
	return p.isAlive
}

func (p *Plant) Mate(e Entity, population []Entity) []Entity {
	return []Entity{}
}

func (p *Plant) Move(others []Entity) {}

func (p *Plant) Pos() Position {
	return p.Position
}

func (p *Plant) IsAtEndOfLife() bool {
	return false
}

func (p *Plant) Die() {
	p.isAlive = false
}
