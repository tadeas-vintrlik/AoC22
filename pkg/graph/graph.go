package graph

type Graph[T comparable, V any] struct {
	Vertices map[T]Vertex[T, V]
}

type Vertex[T comparable, V any] struct {
	Key   T
	Value V
	Edges []T
}

func New[T comparable, V any]() Graph[T, V] {
	g := Graph[T, V]{}
	g.Vertices = make(map[T]Vertex[T, V])
	return g
}

// W is the weight function to be used for the FloydWarshall aglorithm. returns a distance map.
func (g Graph[T, V]) FloydWarshall(w func(u Vertex[T, V], v Vertex[T, V]) int) map[[2]T]int {
	dist := make(map[[2]T]int)

	// Order of map traversal is not guaranteed, therefore we need to first collect the keys
	// into slices and then iterate that slice.
	vertices := []Vertex[T, V]{}
	for _, v := range g.Vertices {
		vertices = append(vertices, v)
	}

	for _, v := range vertices {
		for _, v2 := range vertices {
			dist[[2]T{v.Key, v2.Key}] = w(v, v2)
		}
	}
	for _, k := range vertices {
		for _, i := range vertices {
			for _, j := range vertices {
				if dist[[2]T{i.Key, j.Key}] > dist[[2]T{i.Key, k.Key}]+dist[[2]T{k.Key, j.Key}] {
					dist[[2]T{i.Key, j.Key}] = dist[[2]T{i.Key, k.Key}] + dist[[2]T{k.Key, j.Key}]
				}
			}
		}
	}

	return dist
}
