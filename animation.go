package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const ANIMATION_DELAY float64 = 1

type Animation struct {
	States     []State
	curr       int
	lastUpdate float64
	Animate    bool
}

func (a *Animation) Load(states []State) {
	a.States = states
	a.curr = 0
	a.lastUpdate = rl.GetTime()
}

func (a *Animation) Play() {
	if a.States == nil {
		return
	}
	time := rl.GetTime()
	a.Animate = true
	if time-a.lastUpdate >= ANIMATION_DELAY {
		a.lastUpdate = time
		a.curr = (a.curr + 1) % len(a.States)
	}
}

func (a *Animation) Stop() {
	a.Animate = false
}

func (a *Animation) GetCurrState() State {
	return a.States[a.curr]
}
