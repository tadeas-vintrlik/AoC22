package grid

import (
	"testing"
)

func TestNew(t *testing.T) {
	x := 6
	y := 4
	g := New[int](x, y)
	if len(g.content) != x*y {
		t.Errorf("Expected size of grid: %d != %d", x*y, len(g.content))
	}
	if x != g.SizeX() {
		t.Errorf("Unexpected xsize: %d != %d", x, g.SizeX())
	}
	if y != g.SizeY() {
		t.Errorf("Unexpected xsize: %d != %d", x, g.SizeY())
	}
}

func TestParse(t *testing.T) {
	in := []string{
		"aaa",
		"bbb",
		"ccc",
	}
	g := Parse(in, func(r rune) int { return int(r - 'a') })
	if g.SizeX() != 3 {
		t.Errorf("Unexpected xsize: %d != %d", 3, g.SizeX())
	}
	if g.SizeY() != 3 {
		t.Errorf("Unexpected ysize: %d != %d", 3, g.SizeY())
	}
	for y := 0; y < g.SizeY(); y++ {
		for x := 0; x < g.SizeX(); x++ {
			if g.Get(x, y) != y {
				t.Errorf("Unexpected value in grid(%d, %d): %d != %d", x, y, g.Get(x, y), y)
			}
		}
	}
}

func TestGetNeighbours(t *testing.T) {
	g := New[int](3, 3)
	n := g.GetNeigbhours(1, 1)
	if len(n) != 4 {
		t.Errorf("Unexpected amount of neighbours: %d != %d", len(n), 4)
	}
	exp := map[int]Node[int]{
		0: {Coord: Coord{X: 0, Y: 1}, Value: 0},
		1: {Coord: Coord{X: 2, Y: 1}, Value: 0},
		2: {Coord: Coord{X: 1, Y: 0}, Value: 0},
		3: {Coord: Coord{X: 1, Y: 2}, Value: 0},
	}
	for i, v := range n {
		if exp[i] != v {
			t.Errorf("Unexpected neighbour: %v != %v", exp[i], v)
		}
	}
}

func TestGetNeighboursEdge(t *testing.T) {
	g := New[int](3, 3)
	n := g.GetNeigbhours(0, 0)
	if len(n) != 2 {
		t.Errorf("Unexpected amount of neighbours: %d != %d", len(n), 2)
	}
	exp := map[int]Node[int]{
		0: {Coord: Coord{X: 1, Y: 0}, Value: 0},
		1: {Coord: Coord{X: 0, Y: 1}, Value: 0},
	}
	for i, v := range n {
		if exp[i] != v {
			t.Errorf("Unexpected neighbour: %v != %v", exp[i], v)
		}
	}
}

func TestFind(t *testing.T) {
	g := New[int](3, 3)
	g.Set(1, 1, 10)
	f := g.Find(10)
	if len(f) != 1 {
		t.Errorf("Unexpected ammount of nodes found: %d != %d", len(f), 1)
	}
	exp := Node[int]{Value: 10, Coord: Coord{X: 1, Y: 1}}
	if f[0] != exp {
		t.Errorf("Unexpected node found: %v != %v", f[0], exp)
	}
}

func TestBFS(t *testing.T) {
	in := []string{
		"abc",
		"efg",
		"hij",
	}
	step := 1
	exp := map[int]rune{
		1: 'b',
		2: 'e',
		3: 'c',
		4: 'f',
		5: 'h',
		6: 'g',
		7: 'i',
		8: 'j',
	}
	g := Parse(in, func(r rune) rune { return r })
	root := g.Find('a')[0]
	g.BFS(
		root,
		func(n Node[rune]) []Node[rune] {
			return g.GetNeigbhours(n.X, n.Y)
		},
		func(parent Node[rune], child Node[rune]) {
			if exp[step] != child.Value {
				t.Errorf("Unexpected node: %c != %c", exp[step], child.Value)
			}
			step++
		},
	)
}

func TestFill(t *testing.T) {
	g := New[int](3, 3)
	g.Fill(1)
	for y := 0; y < g.SizeY(); y++ {
		for x := 0; x < g.SizeY(); x++ {
			if g.Get(x, y) != 1 {
				t.Errorf("Unexpected node: %c != %c", g.Get(x, y), 1)
			}
		}
	}
}

func TestFillPath(t *testing.T) {
	g := New[int](3, 3)
	p := Path{
		Coord{X: 0, Y: 0},
		Coord{X: 0, Y: 2},
		Coord{X: 2, Y: 2},
	}
	g.FillPath(p, 1)
	for y := 0; y < 3; y++ {
		if g.Get(0, y) != 1 {
			t.Errorf("Unexpected node: %c != %c", g.Get(0, y), 1)
		}
	}
	for x := 0; x < 3; x++ {
		if g.Get(x, 2) != 1 {
			t.Errorf("Unexpected node: %c != %c", g.Get(0, x), 1)
		}
	}

	unchaged := []Coord{
		{X: 1, Y: 0},
		{X: 2, Y: 0},
		{X: 1, Y: 1},
		{X: 2, Y: 1},
	}
	for _, v := range unchaged {
		if g.Get(v.X, v.Y) != 0 {
			t.Errorf("Unexpected node: %d != %d", g.Get(v.X, v.Y), 0)
		}
	}
}

func TestFillPathPanic(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("Expected fillPath to panic with diagonal input path")
		}
	}()
	g := New[int](3, 3)
	p := Path{
		Coord{X: 0, Y: 0},
		Coord{X: 1, Y: 1},
	}
	g.FillPath(p, 1)
}

func TestDistance(t *testing.T) {
	act := Distance(Coord{X: 8, Y: 7}, Coord{X: 2, Y: 10})
	exp := 9
	if act != exp {
		t.Errorf("%d != %d", act, exp)
	}
}

func TestFillAround(t *testing.T) {
	expstr := []string{
		"..#..",
		".###.",
		"#####",
		".###.",
		"..#..",
	}
	g := New[string](5, 5)
	g.Fill(".")
	g.FillAround(Coord{X: 2, Y: 2}, 2, "#")
	act := g.String()
	exp := Parse(expstr, func(r rune) string { return string(r) }).String()
	if act != exp {
		t.Errorf("\n%s != \n%s", act, exp)
	}
}

func TestFillAroundLarger(t *testing.T) {
	expstr := []string{
		"........#####...............",
		".......#######..............",
		"......#########.............",
		".....###########............",
		"....#############...........",
		"...###############..........",
		"..#################.........",
		".###################........",
		"..#################.........",
		"...###############..........",
		"....#############...........",
		".....###########............",
		"......#########.............",
		".......#######..............",
		"........#####...............",
		".........###................",
		"..........#.................",
	}
	g := New[string](28, 17)
	g.Fill(".")
	g.FillAround(Coord{X: 10, Y: 7}, 9, "#")
	act := g.String()
	exp := Parse(expstr, func(r rune) string { return string(r) }).String()
	if act != exp {
		t.Errorf("\n%s != \n%s", act, exp)
	}
}

func TestNewNonZero(t *testing.T) {
	expstr := []string{
		"123",
		"456",
		"789",
	}
	g := NewNonZero[rune](-1, 1, -1, 1)
	g.Set(-1, -1, '1')
	g.Set(0, -1, '2')
	g.Set(1, -1, '3')
	g.Set(-1, 0, '4')
	g.Set(0, 0, '5')
	g.Set(1, 0, '6')
	g.Set(-1, 1, '7')
	g.Set(0, 1, '8')
	g.Set(1, 1, '9')
	exp := Parse(expstr, func(r rune) rune { return r })
	for i, v := range g.content {
		if g.content[i] != exp.content[i] {
			t.Errorf("Unexpected node: %c", v)
		}
	}
}

func TestString(t *testing.T) {
	expstr := []string{
		"123",
		"456",
		"789",
	}
	g := Parse(expstr, func(r rune) int { return int(r - '0') })
	act := g.String()
	exp := "123\n456\n789\n"
	if act != exp {
		t.Errorf("%s != %s", act, exp)
	}
}

func TestStringNonZero(t *testing.T) {
	g := NewNonZero[int](-1, 1, -1, 1)
	g.Set(-1, -1, 1)
	g.Set(0, -1, 2)
	g.Set(1, -1, 3)
	g.Set(-1, 0, 4)
	g.Set(0, 0, 5)
	g.Set(1, 0, 6)
	g.Set(-1, 1, 7)
	g.Set(0, 1, 8)
	g.Set(1, 1, 9)
	act := g.String()
	exp := "123\n456\n789\n"
	if act != exp {
		t.Errorf("%s != %s", act, exp)
	}
}
