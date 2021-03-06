// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

// Injectors from wire.go:

func InitMission(pname PlayParm, name MosterParm) (Mission, error) {
	player, err := NewPlayer(pname)
	if err != nil {
		return Mission{}, err
	}
	moster, err := NewMoster(name)
	if err != nil {
		return Mission{}, err
	}
	mission, err := NewMission(player, moster)
	if err != nil {
		return Mission{}, err
	}
	return mission, nil
}
