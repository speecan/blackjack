package game

import "testing"

func TestCalcScore(t *testing.T) {
	{
		cs := []*Card{
			NewCard("A"), NewCard("10"),
		}
		expected := uint8(21)
		if s := CalcScore(cs); s != expected {
			t.Error("score was mismatch", s, expected)
		}
	}
	{
		cs := []*Card{
			NewCard("A"), NewCard("A"), NewCard("J"),
		}
		expected := uint8(12)
		if s := CalcScore(cs); s != expected {
			t.Error("score was mismatch", s, expected)
		}
	}
	{
		cs := []*Card{
			NewCard("A"), NewCard("A"), NewCard("7"),
		}
		expected := uint8(19)
		if s := CalcScore(cs); s != expected {
			t.Error("score was mismatch", s, expected)
		}
	}
	{
		cs := []*Card{
			NewCard("A"), NewCard("J"), NewCard("K"),
		}
		expected := uint8(21)
		if s := CalcScore(cs); s != expected {
			t.Error("score was mismatch", s, expected)
		}
	}
	{
		cs := []*Card{
			NewCard("A"), NewCard("J"), NewCard("K"), NewCard("8"),
		}
		expected := uint8(29)
		if s := CalcScore(cs); s != expected {
			t.Error("score was mismatch", s, expected)
		}
	}
}
