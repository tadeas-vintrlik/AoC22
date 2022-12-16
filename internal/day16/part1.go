package day16

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/tadeas-vintrlik/AoC22/pkg/channels"
	"github.com/tadeas-vintrlik/AoC22/pkg/graph"
)

func parseRooms(in <-chan string) graph.Graph[string, int] {
	g := graph.New[string, int]()
	for line := range in {
		r := graph.Vertex[string, int]{}
		re := regexp.MustCompile(`Valve ([A-Z]{2}) has flow rate=(\d+); tunnels? leads? to valves? (.*)`)
		s := re.FindStringSubmatch(line)
		name := s[1]
		flowString := s[2]
		flow, _ := strconv.Atoi(flowString)
		edgestr := s[3]
		r.Value = flow
		edges := strings.Split(edgestr, ", ")
		r.Edges = append(r.Edges, edges...)
		r.Key = name
		g.Vertices[name] = r
	}
	return g
}

func Part1Solver(file string) int {
	g := parseRooms(channels.ReadLines(file))
	dist := g.FloydWarshall(func(u graph.Vertex[string, int], v graph.Vertex[string, int]) int {
		// TODO: this is just a graph BFS
		q := []graph.Vertex[string, int]{}
		q = append(q, u)
		distance := make(map[string]int)
		distance[u.Key] = 0
		for len(q) != 0 {
			c := q[0]
			q = q[1:]
			if c.Key == v.Key {
				return distance[c.Key]
			}
			for _, neighbour := range c.Edges {
				if _, ok := distance[neighbour]; ok {
					continue
				}
				distance[neighbour] = distance[c.Key] + 1
				q = append(q, g.Vertices[neighbour])
			}
		}
		panic("Path not found")
	})

	// Start in root "AA"
	// Find the most viable valve to move to
	// Heuristic is how much presure it could yet release so:
	// (current minutes - distance to possible valve - cost to open) * flow rate of possible valve
	// Move to the most viable valve and open it, minutes - distance to said valve from current
	// Add to pressure release (the heuristic)
	// if all of the possible valves heuristic value is negative we have found our result

	pressure := 0
	current := g.Vertices["AA"]
	minutes := 30
	opened := make(map[string]bool)
	for {
		max := -1
		var next *graph.Vertex[string, int] = nil
		for vname := range g.Vertices {
			if _, ok := opened[vname]; ok {
				continue
			}
			v := g.Vertices[vname]
			heuristic := (minutes - dist[[2]string{current.Key, v.Key}] - 1) * v.Value
			if heuristic > max {
				max = heuristic
				next = &v
			}
		}

		if max <= 0 {
			break
		}

		if max > 0 {
			minutes -= dist[[2]string{current.Key, next.Key}] + 1
			pressure += next.Value * minutes
			opened[next.Key] = true
			current = *next
		}
	}

	return pressure
}
