package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func centerizeRect(r *rl.Rectangle, w, h float32) {
	r.X = (w - r.Width) / 2
	r.Y = (h - r.Height) / 2
}
