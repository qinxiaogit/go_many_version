package foobarbaz

import (
	"fmt"
)

// var SuperSet = wire.NewSet(ProvideFoo, ProvideBar, ProvideBaz)

type Message string

type Monster struct {
	Name string
}

func NewMonster() Monster {
	return Monster{Name: "kitty"}
}

type Player struct {
	Name string
}

func NewPlayer(name string) Player {
	return Player{Name: name}
}

type Mission struct {
	Player  Player
	Monster Monster
}

func NewMission(p Player, m Monster) Mission {
	return Mission{p, m}
}

func (m Mission) Start() {
	fmt.Printf("%s defeats %s, world peace!\n", m.Player.Name, m.Monster.Name)
}

/*
func main() {
	monster := NewMonster()
	player := NewPlayer("dj")
	mission := NewMission(player, monster)

	mission.Start()
}
*/