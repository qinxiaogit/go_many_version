package main

import "github.com/google/wire"
import "wire/foobarbaz"

func InitMission(name string) foobarbaz.Mission {
	wire.Build(foobarbaz.NewMonster, foobarbaz.NewPlayer, foobarbaz.NewMission)
	return foobarbaz.Mission{}
}
