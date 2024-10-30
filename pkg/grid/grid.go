package grid

import (
	"fmt"
	"strings"

	"github.com/tadeas-vintrlik/AoC22/pkg/slice"
	"github.com/tadeas-vintrlik/AoC22/pkg/util"
)

type Grid[T comparable] struct {
	content    []T
	xlen, ylen int
	xmin, ymin int
}

type Node[T any] struct {
	Value T
	Coord
}

type Coord struct {
	X, Y int
}

type Path []Coord

// Create a new empty grid with size of xsize * ysize
func New[T comparable](xsize, ysize int) Grid[T] {
	return Grid[T]{content: make([]T, xsize*ysize), xlen: xsize, ylen: ysize}
}

// Create a new empty grid where the start of x and y need not be zero
func NewNonZero[T comparable](xmin, xmax, ymin, ymax int) Grid[T] {
	xsize := util.Abs(xmin-xmax) + 1
	ysize := util.Abs(ymin-ymax) + 1
	return Grid[T]{content: make([]T, xsize*ysize), xlen: xsize, ylen: ysize, xmin: xmin, ymin: ymin}
}

// Parse a new grid from a grid of runes.
// The transform function can be used to convert it to int for example.
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
	return (y-g.ymin)*g.xlen + (x - g.xmin)
}

// Set value of node on coordinates x and y to v.
func (g *Grid[T]) Set(x, y int, v T) {
	g.content[g.index(x, y)] = v
}

// Get value of node on coordinates x and y.
func (g *Grid[T]) Get(x, y int) T {
	return g.content[g.index(x, y)]
}

// Get neighbours Vertical or Horizontal.
func (g Grid[T]) GetNeigbhours(x, y int) []Node[T] {
	var r []Node[T]
	if x-1 >= g.MinX() {
		r = append(r, Node[T]{Value: g.Get(x-1, y), Coord: Coord{X: x - 1, Y: y}})
	}
	if x+1 < g.MaxX() {
		r = append(r, Node[T]{Value: g.Get(x+1, y), Coord: Coord{X: x + 1, Y: y}})
	}
	if y-1 >= g.MinX() {
		r = append(r, Node[T]{Value: g.Get(x, y-1), Coord: Coord{X: x, Y: y - 1}})
	}
	if y+1 < g.MaxX() {
		r = append(r, Node[T]{Value: g.Get(x, y+1), Coord: Coord{X: x, Y: y + 1}})
	}
	return r
}

// Search the entire grid return list of all nodes with matching value.
func (g Grid[T]) Find(v T) []Node[T] {
	var ret []Node[T]
	for i := g.MinY(); i < g.MaxY(); i++ {
		for j := g.MinX(); j < g.MaxX(); j++ {
			c := g.Get(j, i)
			if c == v {
				ret = append(ret, Node[T]{Value: c, Coord: Coord{X: j, Y: i}})
			}
		}
	}
	return ret
}

// Get size of grid in the x axis.
func (g Grid[T]) SizeX() int {
	return g.xlen
}

// Get size of grid in the y axis.
func (g Grid[T]) SizeY() int {
	return g.ylen
}

func (g Grid[T]) MinX() int {
	return g.xmin
}

func (g Grid[T]) MaxX() int {
	return g.xlen + g.xmin
}

func (g Grid[T]) MinY() int {
	return g.ymin
}

func (g Grid[T]) MaxY() int {
	return g.ylen + g.ymin
}

// Breadth first search on the grid. Useful for example for solving mazes.
//
// The root node is the node where the search starts.
//
// The neighbours function returns all valid neighbours of current node.
// This can be useful for modelling obstacles in the maze for example.
//
// The clb function is called on each node found. The first parameter is the parent.
// The second the node itself.
func (g Grid[T]) BFS(root Node[T], neighbours func(Node[T]) []Node[T], clb func(Node[T], Node[T])) {
	queue := []Node[T]{root}
	visited := make(map[Node[T]]bool)
	visited[root] = true
	for len(queue) != 0 {
		c := queue[0]
		queue = queue[1:]
		for _, v := range neighbours(c) {
			if _, ok := visited[v]; ok {
				continue
			}
			clb(c, v)
			visited[v] = true
			queue = append(queue, v)
		}
	}
}

// Fill the entire grid with value v.
func (g *Grid[T]) Fill(v T) {
	for y := g.MinY(); y < g.MaxY(); y++ {
		for x := g.MinX(); x < g.MaxX(); x++ {
			g.Set(x, y, v)
		}
	}
}

// Fill path with value v, only works for 90 degree angles (vertical/horizontal)
func (g *Grid[T]) FillPath(p Path, v T) {
	for i := 0; i < len(p)-1; i++ {
		xfrom := slice.Min([]int{p[i].X, p[i+1].X})
		xto := slice.Max([]int{p[i].X, p[i+1].X})
		yfrom := slice.Min([]int{p[i].Y, p[i+1].Y})
		yto := slice.Max([]int{p[i].Y, p[i+1].Y})
		if xfrom-xto == 0 {
			for y := yfrom; y <= yto; y++ {
				g.Set(xfrom, y, v)
			}
		} else if yfrom-yto == 0 {
			for x := xfrom; x <= xto; x++ {
				g.Set(x, yfrom, v)
			}
		} else {
			panic("invalid path contains diagonal movement")
		}
	}
}

// Manhattan distance of two coordinates
func Distance(c1 Coord, c2 Coord) int {
	return util.Abs(c1.X-c2.X) + util.Abs(c1.Y-c2.Y)
}

// Fill around (and including) a given coordinate in the given distance d with value v.
func (g *Grid[T]) FillAround(c Coord, d int, v T) {
	dx := -1
	for y := c.Y - d; y <= c.Y+d; y++ {
		// First half it grows, then it shrinks
		if y <= c.Y {
			dx++
		}
		if y > c.Y {
			dx--
		}
		for x := c.X - dx; x <= c.X+dx; x++ {
			if x >= g.MinX() && y >= g.MinY() && x < g.MaxX() && y < g.MaxY() {
				g.Set(x, y, v)
			}
		}
	}
}

func (g Grid[T]) String() string {
	var sb strings.Builder
	for y := g.MinY(); y < g.MaxY(); y++ {
		for x := g.MinX(); x < g.MaxX(); x++ {
			sb.WriteString(fmt.Sprintf("%v", (g.Get(x, y))))
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}
