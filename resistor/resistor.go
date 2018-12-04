package resistor

type Resistor struct {
	Piece int
	Value float32
}

func Addresistor(p int, v float32) Resistor {
	return Resistor{Piece: p, Value: v}
}
