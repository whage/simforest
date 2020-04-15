package simforest

type Tree struct {
	Position
}

func (t *Tree) Act(population []Entity) []Entity {
	return []Entity{}
}

func (t *Tree) Gender() Gender {
	return Female // All Trees are Female if I say so!
}

func (t *Tree) IsAdult() bool {
	return false // Trees are always young.
}

func (t *Tree) IsAlive() bool {
	return true // And live forever
}

func (t *Tree) Mate(e Entity, population []Entity) []Entity {
	return []Entity{}
}

func (t *Tree) Move(others []Entity) {}

func (t *Tree) Pos() Position {
	return t.Position
}

func (t *Tree) Render() Marker {
	return Marker{
		DarkGreen,
		"T",
	}
}

func (t *Tree) IsAtEndOfLife() bool {
	return false
}

func (t *Tree) Die() {
}
