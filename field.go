package main

import (
	// "fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

type Field struct {
	Bound   rl.Rectangle
	Col     rl.Color
	Size    int
	Content [][]byte
}

func NewField() Field {
	var size float32 = 400
	f := Field{
		Bound: rl.Rectangle{
			Width:  size,
			Height: size,
		},
		Col: rl.Color{
			R: 0x18,
			G: 0x18,
			B: 0x18,
			A: 0xFF,
		},
		Size: 4,
	}
	f.Content = make([][]byte, f.Size)
	for i, _ := range f.Content {
		f.Content[i] = make([]byte, f.Size)
	}
	f.genBaseContent()
	centerizeRect(&f.Bound, WDT, HGT)
	return f
}

func (f *Field) genBaseContent() {
	for i := 0; i < f.Size; i++ {
		for j := 0; j < f.Size; j++ {
			f.Content[i][j] = byte(j)
		}
	}
}

func (f *Field) Update() {
	f.updateBtns()
}

func (f *Field) updateBtns() {
	mouse := rl.GetMousePosition()
	var x, y, sz, padding float32
	var r rl.Rectangle

	padding = 10
	sz = f.Bound.Width / float32(f.Size)

	col := rl.Red
	col.A = 0x80

	x = f.Bound.X
	y = f.Bound.Y - sz
	for i := 0; i < f.Size; i++ {
		r.X = x + padding/2
		r.Y = y + padding/2
		r.Width = sz - padding
		r.Height = sz - padding
		if rl.CheckCollisionPointRec(mouse, r) && rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			f.moveCol(i)
		}
		x += sz
	}
	x = f.Bound.X - sz
	y = f.Bound.Y
	for i := 0; i < f.Size; i++ {
		r.X = x + padding/2
		r.Y = y + padding/2
		r.Width = sz - padding
		r.Height = sz - padding
		if rl.CheckCollisionPointRec(mouse, r) && rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			f.moveRow(i)
		}
		y += sz
	}
}

func (f *Field) moveCol(idx int) {
	fst := f.Content[0][idx]
	for j := 0; j < f.Size-1; j++ {
		f.Content[j][idx] = f.Content[j+1][idx]
	}
	f.Content[f.Size-1][idx] = fst
}

func (f *Field) moveRow(idx int) {
	fst := f.Content[idx][0]
	for j := 0; j < f.Size-1; j++ {
		f.Content[idx][j] = f.Content[idx][j+1]
	}
	f.Content[idx][f.Size-1] = fst
}

func (f *Field) Draw() {
	rl.DrawRectangleRounded(f.Bound, 0.1, 0, f.Col)
	f.drawContent()
	f.drawBtns()
}

func (f *Field) drawContent() {
	var sz, x, y float32
	sz = f.Bound.Width / float32(f.Size)
	x = f.Bound.X
	y = f.Bound.Y

	for i := 0; i < f.Size; i++ {
		for j := 0; j < f.Size; j++ {
			rl.DrawCircle(int32(x+sz/2), int32(y+sz/2), sz/2*0.9, getColByVal(f.Content[i][j]))
			x += sz
		}
		x = f.Bound.X
		y += sz
	}
}

func (f *Field) drawBtns() {
	var padding, sz, x, y float32

	padding = 10
	sz = f.Bound.Width / float32(f.Size)
	x = f.Bound.X
	y = f.Bound.Y - sz

	for i := 0; i < f.Size; i++ {
		drawArrow(x, y, sz, padding, Up, rl.White)
		x += sz
	}
	x = f.Bound.X - sz
	y = f.Bound.Y
	for i := 0; i < f.Size; i++ {
		drawArrow(x, y, sz, padding, Left, rl.White)
		y += sz
	}
}

func drawArrow(x, y, sz, padding float32, dir Direction, col rl.Color) {
	var thickness float32 = 4
	var start, vDir, v1, v2 rl.Vector2
	switch dir {
	case Up:
		start.X = x + sz/2
		start.Y = y + sz
		vDir.X = 0
		vDir.Y = -sz / 2
	case Down:
		start.X = x + sz/2
		start.Y = y
		vDir.X = 0
		vDir.Y = sz / 2
	case Left:
		start.X = x + sz
		start.Y = y + sz/2
		vDir.X = -sz / 2
		vDir.Y = 0
	case Right:
		start.X = x
		start.Y = y + sz/2
		vDir.X = sz / 2
		vDir.Y = 0
	}

	var angle float32 = rl.Pi/2 + rl.Pi/4
	v1 = rl.Vector2Rotate(vDir, angle)
	v2 = rl.Vector2Rotate(vDir, -angle)

	v1 = rl.Vector2Add(rl.Vector2Add(start, vDir), v1)
	v2 = rl.Vector2Add(rl.Vector2Add(start, vDir), v2)
	start = rl.Vector2Add(start, vDir)

	rl.DrawLineEx(start, v1, thickness, col)
	rl.DrawLineEx(start, v2, thickness, col)
}

func getColByVal(col byte) rl.Color {
	switch col {
	case 0:
		return rl.Color{
			R: 0xFF,
			G: 0x00,
			B: 0x00,
			A: 0xFF,
		}
	case 1:
		return rl.Color{
			R: 0x00,
			G: 0xFF,
			B: 0x00,
			A: 0xFF,
		}
	case 2:
		return rl.Color{
			R: 0x00,
			G: 0x00,
			B: 0xFF,
			A: 0xFF,
		}
	case 3:
		return rl.Color{
			R: 0x80,
			G: 0x00,
			B: 0x80,
			A: 0xFF,
		}
	}
	return rl.Color{}
}
