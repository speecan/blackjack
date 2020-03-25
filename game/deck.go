package game

import (
	"math/rand"
	"time"
)

type (
	Deck struct {
		cards  []*Card
		Drawed []*Card
	}
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func NewDeck() *Deck {
	d := &Deck{
		cards: make([]*Card, 0),
	}
	captions := []string{"A", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	for _, n := range captions {
		for i := 0; i < 4; i++ {
			if c := NewCard(n); c != nil {
				d.cards = append(d.cards, c)
			}
		}
	}
	d.Shuffle()
	return d
}

func (d *Deck) Shuffle() {
	cs := d.cards
	for i := range cs {
		j := rand.Intn(i + 1)
		cs[i], cs[j] = cs[j], cs[i]
	}
	d.cards = cs
}

func (d *Deck) Draw() *Card {
	if len(d.cards) == 0 {
		return nil
	}
	c := d.cards[0]
	d.cards = d.cards[1:]
	d.Drawed = append(d.Drawed, c)
	return c
}
