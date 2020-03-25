package game

type (
	Card struct {
		Number  uint8
		Caption string
		Point   uint8
	}
)

func NewCard(c string) *Card {
	switch c {
	case "A":
		return &Card{Number: 1, Caption: c, Point: 1}
	case "2":
		return &Card{Number: 2, Caption: c, Point: 2}
	case "3":
		return &Card{Number: 3, Caption: c, Point: 3}
	case "4":
		return &Card{Number: 4, Caption: c, Point: 4}
	case "5":
		return &Card{Number: 5, Caption: c, Point: 5}
	case "6":
		return &Card{Number: 6, Caption: c, Point: 6}
	case "7":
		return &Card{Number: 7, Caption: c, Point: 7}
	case "8":
		return &Card{Number: 8, Caption: c, Point: 8}
	case "9":
		return &Card{Number: 9, Caption: c, Point: 9}
	case "10":
		return &Card{Number: 10, Caption: c, Point: 10}
	case "J":
		return &Card{Number: 11, Caption: c, Point: 10}
	case "Q":
		return &Card{Number: 12, Caption: c, Point: 10}
	case "K":
		return &Card{Number: 13, Caption: c, Point: 10}
	}
	return nil
}

func CalcScore(cs []*Card) uint8 {
	var s uint8
	aceNumber := 0
	for _, c := range cs {
		s += c.Point
		if c.Number == 1 {
			aceNumber++
		}
	}
	for i := 0; i < aceNumber; i++ {
		if s+10 <= 21 {
			s += 10
		}
	}
	return s
}

func (x *Card) String() string {
	return x.Caption
}
