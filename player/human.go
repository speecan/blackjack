package player

import (
	"fmt"
	"os"
	"strings"

	"github.com/speecan/blackjack/game"
)

type (
	Human struct {
		Name  string
		cards []*game.Card
	}
)

func (x *Human) CallHit(drawed []*game.Card) bool {
	var input string
	fmt.Print("カードを引きますか? (y/n): ")
	fmt.Fscanln(os.Stdin, &input)
	input = strings.ToLower(input)
	if input == "y" {
		return true
	} else if input == "n" {
		return false
	}
	fmt.Println("入力をやり直して下さい")
	return x.CallHit(drawed)
}

func (x *Human) GetName() string {
	return x.Name
}
