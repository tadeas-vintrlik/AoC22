package tree

import "testing"

type person struct {
	name   string
	active bool
}

func createTestTree() *Tree[person] {
	t := &Tree[person]{Value: person{"F", true}}
	Append(t, person{"B", true})
	Append(t, person{"G", true})
	Append(t.children[0], person{"A", false})
	Append(t.children[0], person{"D", true})
	Append(t.children[1], person{"I", false})
	return t
}

func TestAppend(t *testing.T) {
	d := &Tree[person]{}
	if len(d.children) != 0 {
		t.Errorf("invalid empty tree")
	}
	n := person{"Jane", false}
	Append(d, n)
	if len(d.children) != 1 {
		t.Errorf("append does not add new children")
	}
	if d.children[0].Value != n {
		t.Errorf("append added different value")
	}
}

func TestChildren(t *testing.T) {
	d := createTestTree()
	s := Children(d)
	if len(s) != 2 {
		t.Errorf("children does not return correct amount of children")
	}
}

func TestPostorderCollect(t *testing.T) {
	d := createTestTree()
	var all []person
	PostorderCollect(d, &all)
	exp := []person{
		{"A", false},
		{"D", true},
		{"B", true},
		{"I", false},
		{"G", true},
		{"F", true},
	}
	if len(all) != len(exp) {
		t.Errorf("unexpected length of slice from PostorderCoolect actual: %d, expected %d", len(all), len(exp))
	}
	for i := range all {
		if all[i] != exp[i] {
			t.Errorf("unexpected element in slice, actual: %v, expected: %v", all[i], exp[i])
		}
	}
}

func TestPostorderCollectFilter(t *testing.T) {
	d := createTestTree()
	var all []person
	PostorderCollectFilter(d, func(p person) bool {
		return !p.active
	}, &all)
	exp := []person{
		{"A", false},
		{"I", false},
	}
	if len(all) != len(exp) {
		t.Errorf("unexpected length of slice from PostorderCoolect actual: %d, expected %d", len(all), len(exp))
	}
	for i := range all {
		if all[i] != exp[i] {
			t.Errorf("unexpected element in slice, actual: %v, expected: %v", all[i], exp[i])
		}
	}
}

func TestPostorderCallback(t *testing.T) {
	d := createTestTree()
	PostorderCallback(d, func(t *Tree[person]) {
		t.Value.active = !t.Value.active
	})
	var all []person
	PostorderCollect(d, &all)
	exp := []person{
		{"A", true},
		{"D", false},
		{"B", false},
		{"I", true},
		{"G", false},
		{"F", false},
	}
	if len(all) != len(exp) {
		t.Errorf("unexpected length of slice from PostorderCollect actual: %d, expected %d", len(all), len(exp))
	}
	for i := range all {
		if all[i] != exp[i] {
			t.Errorf("unexpected element in slice, actual: %v, expected: %v", all[i], exp[i])
		}
	}
}
