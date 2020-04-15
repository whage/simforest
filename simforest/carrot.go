package simforest

type Carrot struct {
	Plant
}

func (p *Plant) Render() Marker {
	return Marker{
		Orange,
		"c",
	}
}
