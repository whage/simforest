package simforest

type Population []Creature

func (p Population) FilterBunnies() []Bunny {
	var results []Bunny
	for i, _ := range p {
		v, ok := p[i].(*Bunny)
		if ok {
			results = append(results, *v)
		}
	}
	return results

}
