package main

import "fmt"

type Moster struct {
	Name string
}

type MosterParm string

func NewMoster(name MosterParm) (Moster, error) {
	return Moster{Name: string(name)}, nil
}

type Player struct {
	Name string
}

type PlayParm string

func NewPlayer(name PlayParm) (Player, error) {
	return Player{Name: string(name)}, nil
}

type Mission struct {
	Player Player
	Moster Moster
}

func NewMission(p Player, m Moster) (Mission, error) {
	return Mission{p, m}, nil
}

func (m Mission) start() {
	fmt.Printf("%s defeats %s, world peace", m.Player.Name, m.Moster.Name)
}

func main() {
	Mission, err := InitMission("dj", "kitty")
	if err != nil {
		fmt.Println(err)
	}
	Mission.start()

}
