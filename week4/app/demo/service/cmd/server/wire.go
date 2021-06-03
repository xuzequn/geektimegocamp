// +build wireinject

//+build wireinject

package main

import "github.com/google/wire"

func InitMission(pname PlayParm, name MosterParm) (Mission, error) {
	wire.Build(NewMoster, NewPlayer, NewMission)
	return Mission{}, nil
}
