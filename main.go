package main

import (
	"fmt"
	"strings"

	"github.com/speecan/blackjack/game"
	"github.com/speecan/blackjack/player"
)

type (
	Game struct {
		deck     *game.Deck
		statuses []*status
	}
	status struct {
		player     player.Player
		hands      []*game.Card
		isFinished bool
	}
)

func NewGame(ps []player.Player) *Game {
	d := game.NewDeck()
	ss := make([]*status, 0)
	for _, p := range ps {
		s := &status{
			player: p,
			hands:  []*game.Card{d.Draw(), d.Draw()},
		}
		ss = append(ss, s)
	}
	return &Game{
		deck:     d,
		statuses: ss,
	}
}

func (g *Game) next() bool {
	for _, s := range g.statuses {
		if !s.isFinished {
			return true
		}
	}
	return false
}

func (g *Game) String() string {
	res := []string{}
	for _, s := range g.statuses {
		cs := []string{}
		for _, c := range s.hands {
			cs = append(cs, c.String())
		}
		finished := ""
		if s.isFinished {
			finished = "(stop)"
		}
		res = append(res, fmt.Sprintf("%s (%d): %s %s", s.player.GetName(), game.CalcScore(s.hands), strings.Join(cs, ", "), finished))
	}
	return strings.Join(res, "\n")
}

func (g *Game) Winner() player.Player {
	w := -1
	r := -1000
	for i, s := range g.statuses {
		c := game.CalcScore(s.hands)
		if c > 21 {
			continue
		}
		if int(c) > r {
			r = int(c)
			w = i
		} else if int(c) == r {
			alen := len(g.statuses[w].hands)
			blen := len(s.hands)
			if alen == blen { // 同じ点数かつ枚数の同じ場合は r をそのままにして winner のみ nil にする
				w = -1
			}
			if blen < alen { // 同じ点数の場合は枚数の少ないほうが勝ち
				r = int(c)
				w = i
			}
		}
	}
	if w == -1 {
		return nil
	}
	return g.statuses[w].player
}

func main() {
	players := []player.Player{
		&player.Human{Name: "A"},
		&player.Human{Name: "B"},
	}
	w := startGame(players)
	if w != nil {
		fmt.Println("勝者:", w.GetName())
	}
}

func startGame(ps []player.Player) (winner player.Player) {
	g := NewGame(ps)
	for g.next() {
		fmt.Println(g.String())
		for _, s := range g.statuses {
			if s.isFinished {
				continue
			}
			fmt.Println(s.player.GetName(), "の番です")
			if s.player.CallHit(s.hands) {
				c := g.deck.Draw()
				fmt.Println(s.player.GetName(), "は", c.String(), "を引きました")
				s.hands = append(s.hands, c)
				if game.CalcScore(s.hands) > 21 {
					fmt.Println(s.player.GetName(), ": burst", game.CalcScore(s.hands))
					s.isFinished = true
				}
			} else {
				s.isFinished = true
			}
		}
	}
	fmt.Println("ゲーム終了")
	fmt.Println(g.String())
	winner = g.Winner()
	return
}
