package main

import (
	"math/rand"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

type Content [][]byte

type Field struct {
	Bound   rl.Rectangle
	Col     rl.Color
	Size    int
	Content Content
}

func NewField() Field {
	var size float32 = 400
	f := Field{
		Bound: rl.Rectangle{
			Width:  size,
			Height: size,
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
			f.Content.moveCol(i)
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
			f.Content.moveRow(i)
		}
		y += sz
	}
}

func (f *Field) MoveRandomly(num int) {
	if num <= 0 {
		return
	}
	size := len(f.Content)
	f.Content.moveRow(rand.Int() % size)
	num -= 1
	for i := 0; i < num; i++ {
		if rand.Int() % 2 == 0 {
			f.Content.moveRow(rand.Int() % size)
		} else {
			f.Content.moveCol(rand.Int() % size)
		}
	}
}

func (c Content) moveCol(idx int) {
	fst := c[0][idx]
	for i := 0; i < len(c)-1; i++ {
		c[i][idx] = c[i+1][idx]
	}
	c[len(c)-1][idx] = fst
}

func (c Content) moveRow(idx int) {
	fst := c[idx][0]
	for i := 0; i < len(c)-1; i++ {
		c[idx][i] = c[idx][i+1]
	}
	c[idx][len(c)-1] = fst
}

func (c Content) moveColReverse(idx int) {
	lst := c[len(c)-1][idx]
	for i := len(c) - 1; i > 0; i-- {
		c[i][idx] = c[i-1][idx]
	}
	c[0][idx] = lst
}

func (c Content) moveRowReverse(idx int) {
	lst := c[idx][len(c)-1]
	for i := len(c) - 1; i > 0; i-- {
		c[idx][i] = c[idx][i-1]
	}
	c[idx][0] = lst
}

func (f *Field) Draw() {
	rl.DrawRectangleRounded(f.Bound, 0.1, 0, COL_BOX_BACKGROUND)
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
		drawArrow(x, y, sz, padding, Up, COL_ARROW)
		x += sz
	}
	x = f.Bound.X - sz
	y = f.Bound.Y
	for i := 0; i < f.Size; i++ {
		drawArrow(x, y, sz, padding, Left, COL_ARROW)
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
		return COL_CIRC1
	case 1:
		return COL_CIRC2
	case 2:
		return COL_CIRC3
	case 3:
		return COL_CIRC4
	}
	return rl.Color{}
}
