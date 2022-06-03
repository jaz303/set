package set

import "testing"

func TestOf(t *testing.T) {
	s := Of(1, 2, 3)
	if s.Size() != 3 || !s.ContainsSlice([]int{1, 2, 3}) {
		t.Fail()
	}
}

func TestNewSetIsEmpty(t *testing.T) {
	s := make(Set[int])
	if !s.Empty() {
		t.Fail()
	}
}

func TestSetWithItemsIsNotEmpty(t *testing.T) {
	s := make(Set[int])
	s.Add(1)
	if s.Empty() {
		t.Fail()
	}
}

func TestNewSetHasZeroSize(t *testing.T) {
	s := make(Set[int])
	if s.Size() != 0 {
		t.Fail()
	}
}

func TestNewSetReportsCorrectSize(t *testing.T) {
	s := make(Set[int])
	s.Add(1)
	s.Add(1)
	s.Add(2)
	s.Add(3)
	s.Remove(2)
	if s.Size() != 2 {
		t.Fail()
	}
}

func TestContainsReturnsFalseWhenNoItem(t *testing.T) {
	s := make(Set[int])
	if s.Contains(1) {
		t.Fail()
	}
}

func TestContainsReturnsTrueWhenItemExists(t *testing.T) {
	s := make(Set[int])
	s.Add(100)
	if !s.Contains(100) {
		t.Fail()
	}
}

func TestContainsSliceReturnsFalseWhenNotAllItemsPresent(t *testing.T) {
	s := make(Set[int])
	s.Add(1)
	s.Add(2)
	if s.ContainsSlice([]int{2, 3}) {
		t.Fail()
	}
}

func TestContainsSliceReturnsTrueWhenAllItemsPresent(t *testing.T) {
	s := make(Set[int])
	s.Add(1)
	s.Add(2)
	if !s.ContainsSlice([]int{1, 2}) {
		t.Fail()
	}
}

func TestClearRemovesAllItems(t *testing.T) {
	s := make(Set[int])
	s.Add(1)
	s.Add(2)
	s.Add(3)
	s.Clear()
	if len(s) != 0 || !s.Empty() || s.Size() != 0 {
		t.Fail()
	}
}

func TestAdd(t *testing.T) {
	s := make(Set[int])
	s.Add(100)
	_, present := s[100]
	if !present || len(s) != 1 {
		t.Fail()
	}
}

func TestAddSlice(t *testing.T) {
	s := make(Set[int])
	s.AddSlice([]int{1, 2, 3})

	_, p1 := s[1]
	_, p2 := s[2]
	_, p3 := s[3]

	if len(s) != 3 || !p1 || !p2 || !p3 {
		t.Fail()
	}
}

func TestAddSet(t *testing.T) {
	s := make(Set[int])
	s.AddSet(Of(1, 2, 3))

	_, p1 := s[1]
	_, p2 := s[2]
	_, p3 := s[3]

	if len(s) != 3 || !p1 || !p2 || !p3 {
		t.Fail()
	}
}

func TestRemove(t *testing.T) {
	s := make(Set[int])
	s.Add(100)
	s.Remove(100)
	_, present := s[100]
	if present || len(s) != 0 {
		t.Fail()
	}
}

func TestRemoveSlice(t *testing.T) {
	s := Of(1, 2, 3, 4, 5)
	s.RemoveSlice([]int{2, 3, 4})

	_, p1 := s[1]
	_, p5 := s[5]

	if len(s) != 2 || !p1 || !p5 {
		t.Fail()
	}
}

func TestRemoveSet(t *testing.T) {
	s := Of(1, 2, 3, 4, 5)
	s.RemoveSet(Of(2, 3, 4))

	_, p1 := s[1]
	_, p5 := s[5]

	if len(s) != 2 || !p1 || !p5 {
		t.Fail()
	}
}

func TestItems(t *testing.T) {
	s := Of(1, 2, 3)
	items := s.Items()

	contains := func(i int) bool {
		for _, v := range items {
			if v == i {
				return true
			}
		}
		return false
	}

	if len(items) != 3 || !contains(1) || !contains(2) || !contains(3) {
		t.Fail()
	}
}

func TestUnion(t *testing.T) {
	a := Of(1, 2, 3)
	b := Of(3, 4, 5)
	u := Union(a, b)

	if u.Size() != 5 || !u.ContainsSlice([]int{1, 2, 3, 4, 5}) {
		t.Fail()
	}
}

func TestIntersection(t *testing.T) {
	a := Of(1, 2, 3)
	b := Of(3, 4, 5)
	i := Intersection(a, b)

	if i.Size() != 1 || !i.Contains(3) {
		t.Fail()
	}
}

func TestDifference(t *testing.T) {
	a := Of(1, 2, 3, 4)
	b := Of(2, 4, 5)
	d := Difference(a, b)

	if d.Size() != 2 || !d.ContainsSlice([]int{1, 3}) {
		t.Fail()
	}
}
