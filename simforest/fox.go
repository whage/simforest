package simforest

type Fox struct {
	Animal
}

func (f *Fox) Move() {
	randomStep(&f.pos, f.environment)
}

func (f *Fox) Pos() Position {
	return f.pos
}

func (f *Fox) IncreaseAge() {
	f.age += 1
}

func (f *Fox) TryToMate(others []Breeder) Breeder {
	if (f.gender == Male) {
		for _, o := range others {
			o, ok := o.(*Fox)
			if ok {
				if f.Pos().IsNearby(o.Pos()) && (o.gender == Female) && f.age > 30 && o.age > 30 {
					return CreateBunny(Position{f.pos.X, f.pos.Y}, f.environment)
				}
			}
		}
	}
	return nil
}
