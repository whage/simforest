package simforest

type Carrot struct {
	Position
}

func (c *Carrot) Act(population []Entity) []Entity {
	return []Entity{}
}

func (c *Carrot) Gender() Gender {
	return Female // All Carrots are Female if I say so!
}

func (c *Carrot) IsAdult() bool {
	return false // Carrots are always young.
}

func (c *Carrot) IsAlive() bool {
	return true // And live forever
}

func (c *Carrot) Mate(e Entity, population []Entity) []Entity {
	return []Entity{}
}

func (c *Carrot) Move(others []Entity) {}

func (c *Carrot) Pos() Position {
	return c.Position
}

func (c *Carrot) Render() Marker {
	return Marker{
		Orange,
		"c",
	}
}
