package grid

// TODO: Grid tests

type Grid[T comparable] struct {
	content    []T
	xlen, ylen int
}

func New[T comparable](xsize, ysize int) Grid[T] {
	return Grid[T]{content: make([]T, xsize*ysize), xlen: xsize, ylen: ysize}
}

func Parse[T comparable](lines []string, transform func(rune) T) Grid[T] {
	xlen := len(lines[0])
	ylen := len(lines)
	g := Grid[T]{xlen: xlen, ylen: ylen, content: make([]T, xlen*ylen)}
	for i, line := range lines {
		for j, v := range line {
			g.Set(j, i, transform(v))
		}
	}
	return g
}

func (g Grid[T]) index(x, y int) int {
	return y*g.xlen + x
}

func (g *Grid[T]) Set(x, y int, v T) {
	g.content[g.index(x, y)] = v
}

func (g *Grid[T]) Get(x, y int) T {
	return g.content[y*g.xlen+x]
}

type Node[T any] struct {
	Value T
	X, Y  int
}

// Get neighbours Vertical or Horizontal
func (g Grid[T]) GetNeigbhours(x, y int) []Node[T] {
	var r []Node[T]
	if x-1 >= 0 {
		r = append(r, Node[T]{Value: g.Get(x-1, y), X: x - 1, Y: y})
	}
	if x+1 < g.xlen {
		r = append(r, Node[T]{Value: g.Get(x+1, y), X: x + 1, Y: y})
	}
	if y-1 >= 0 {
		r = append(r, Node[T]{Value: g.Get(x, y-1), X: x, Y: y - 1})
	}
	if y+1 < g.ylen {
		r = append(r, Node[T]{Value: g.Get(x, y+1), X: x, Y: y + 1})
	}
	return r
}

func (g Grid[T]) Find(v T) []Node[T] {
	var ret []Node[T]
	for i := 0; i < g.ylen; i++ {
		for j := 0; j < g.xlen; j++ {
			c := g.Get(j, i)
			if c == v {
				ret = append(ret, Node[T]{Value: c, X: j, Y: i})
			}
		}
	}
	return ret
}

func (g Grid[T]) SizeX() int {
	return g.xlen
}

func (g Grid[T]) SizeY() int {
	return g.ylen
}
