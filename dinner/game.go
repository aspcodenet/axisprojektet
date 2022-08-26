package dinner

import (
	"fmt"
	"math/rand"
	"time"
)

type Human struct {
	Namn  string
	Level int
}

type Fly struct {
	Namn  string
	Level int
}

type IActionable interface {
	act()
}

type ILevelCalculator interface {
	maybeLevelUp()
}

type IActionableLevelCalculator interface {
	IActionable
	ILevelCalculator
}

type IGameObject interface {
	maybeLevelUp()
	act()
	/*
		...
		...
		...*/
}

func (t *Fly) maybeLevelUp() {
}

func (t *Human) maybeLevelUp() {
	t.Level++
}

func (t *Fly) act() {
	var actions = [...]string{"flyger", "surrar'", "landar i maten"}
	var action = actions[rand.Intn(len(actions))]
	fmt.Printf("%s %s\n", t.Namn, action)
}

func (t *Human) act() {
	var actions = [...]string{"rapar", "smaskar", "tänker", "pratar"}
	var action = actions[rand.Intn(len(actions))]
	fmt.Printf("%s %s\n", t.Namn, action)
}

func gameMove(guest IGameObject) {
	guest.act()
	guest.maybeLevelUp()
}

func Run() {
	var allGuests []IGameObject
	p := &Human{Namn: "Stefan"}
	p2 := &Human{Namn: "Kerstin"}
	fly := &Fly{Namn: "Flugan"}

	allGuests = append(allGuests, p, p2, fly)

	for {
		for _, guest := range allGuests {
			gameMove(guest)
		}
		time.Sleep(time.Millisecond * 1500)
	}

}
