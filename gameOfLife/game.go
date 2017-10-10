package main

import (
	"fmt"
	"time"
	"math/rand"
	"bytes"
)

type Field struct {
	array         [][]bool
	width, height int
}

type Life struct {
	a, b          *Field
	width, height int
}

// NewLife returns a new Life game state with a random initial state.
func NewLife(width, height int) *Life {
	a := NewField(width, height)
	for i := 0; i < (width * height / 4); i++ {
		a.SetValue(rand.Intn(width), rand.Intn(height), true)
	}
	return &Life{
		a:     a, b: NewField(width, height),
		width: width, height: height,
	}
}

func (field *Field) IsPopulated(x int, y int) bool {
	x += field.width
	x %= field.width
	y += field.height
	y %= field.height
	return field.array[x][y]
}

func (field *Field) SetValue(x, y int, value bool) {
	field.array[x][y] = value
}

func NewField(width, height int) *Field {
	array := make([][]bool, width)
	for i := range array {
		array[i] = make([]bool, height)
	}
	return &Field{array, width, height}
}

// Next returns the state of the specified cell at the next time step.
func (field *Field) Next(x, y int) bool {
	// Count the adjacent cells that are IsPopulated.
	isPopulated := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if (j != 0 || i != 0) && field.IsPopulated(x+i, y+j) {
				isPopulated++
			}
		}
	}
	// Return next state according to the game rules:
	//   exactly 3 neighbors: on,
	//   exactly 2 neighbors: maintain current state,
	//   otherwise: off.
	return isPopulated == 3 || isPopulated == 2 && field.IsPopulated(x, y)
}

// Step advances the game by one instant, recomputing and updating all cells.
func (life *Life) Step() {
	// Update the state of the next field (b) from the current field (a).
	for y := 0; y < life.height; y++ {
		for x := 0; x < life.width; x++ {
			life.b.SetValue(x, y, life.a.Next(x, y))
		}
	}
	// Swap fields a and b.
	life.a, life.b = life.b, life.a
}

// String returns the game board as a string.
func (life *Life) String() string {
	var buf bytes.Buffer
	for y := 0; y < life.height; y++ {
		for x := 0; x < life.width; x++ {
			b := byte(' ')
			if life.a.IsPopulated(x, y) {
				b = '*'
			}
			buf.WriteByte(b)
		}
		buf.WriteByte('\n')
	}
	return buf.String()
}

func main()  {
	l := NewLife(40, 15)
	for i := 0; i < 300; i++ {
		l.Step()
		fmt.Print("\x0c", l) // Clear screen and print field.
		time.Sleep(time.Second / 50)
	}
}
