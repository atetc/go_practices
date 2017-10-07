package gameOfLife

type Field struct {
	s [][]bool
	w, h int
}

func main()  {
	field := newField(3, 3)
	field.set(1, 1, true)
	s := field.alive(-1, 0)
}
func (field *Field) alive(x int, y int) bool {
	return field.s[x][y]
}

func (field *Field) set(x, y int, value bool) {
	field.s[x][y] = value
}

func newField(w, h int) *Field {
	bools := make([][]bool, w)
	for i := range bools {
		bools[i] = make([]bool, h)
	}
	return &Field{bools, w, h}
}
