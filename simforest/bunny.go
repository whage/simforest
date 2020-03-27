package simforest

type Bunny struct {
	Animal
}

func (b *Bunny) Move() {
	randomStep(&b.pos, b.environment)
}

func (b *Bunny) Pos() Position {
	return b.pos
}

func (b *Bunny) IncreaseAge() {
	b.age += 1
}

func (b *Bunny) TryToMate(others []Breeder) Breeder {
	if (b.gender == Male) {
		for _, o := range others {
			o, ok := o.(*Bunny)
			if ok {
				if b.Pos().IsNearby(o.Pos()) && (o.gender == Female) && b.age > 30 && o.age > 30 {
					return CreateBunny(Position{b.pos.X, b.pos.Y}, b.environment)
				}
			}
		}
	}
	return nil
}
