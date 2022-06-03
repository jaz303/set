// Package set provides a generic set implementation
package set

import (
	"encoding/json"
)

// Set implements a set - an unordered collection of items wherein each item is
// unique.
//
// The underlying storage mechanism is a map of set items to struct{}. To create
// a set, use make():
//
//    set := make(Set[int])
type Set[T comparable] map[T]struct{}

// Of returns a set comprising the specified items.
func Of[T comparable](vs ...T) Set[T] {
	out := make(Set[T])
	for _, v := range vs {
		out[v] = struct{}{}
	}
	return out
}

// Empty returns true if the set is empty, false otherwise.
func (s Set[T]) Empty() bool { return len(s) == 0 }

// Size returns the number of items in the set.
func (s Set[T]) Size() int { return len(s) }

// Contains returns true if the set contains v, false otherwise.
func (s Set[T]) Contains(v T) bool { _, ok := s[v]; return ok }

// ContainsSlice returns true is the set contains all elements of vs, false otherwise.
func (s Set[T]) ContainsSlice(vs []T) bool {
	for _, v := range vs {
		if _, ok := s[v]; !ok {
			return false
		}
	}
	return true
}

// Clear removes all items from the set.
func (s Set[T]) Clear() {
	for k := range s {
		delete(s, k)
	}
}

// Add adds v to the set.
func (s Set[T]) Add(v T) {
	s[v] = struct{}{}
}

// AddSlice adds all elements of vs to the set.
func (s Set[T]) AddSlice(vs []T) {
	for _, v := range vs {
		s[v] = struct{}{}
	}
}

// AddSet adds all items in set vs to the set.
func (s Set[T]) AddSet(vs Set[T]) {
	for v := range vs {
		s[v] = struct{}{}
	}
}

// Remove removes v from the set.
func (s Set[T]) Remove(v T) {
	delete(s, v)
}

// RemoveSlice removes all elements of vs from the set.
func (s Set[T]) RemoveSlice(vs []T) {
	for _, v := range vs {
		delete(s, v)
	}
}

// RemoveSet removes all items in vs from the set.
func (s Set[T]) RemoveSet(vs Set[T]) {
	for v := range vs {
		delete(s, v)
	}
}

// Items returns a slice of all items in the set.
func (s Set[T]) Items() []T {
	out := make([]T, 0, len(s))
	for k := range s {
		out = append(out, k)
	}
	return out
}

// Union returns a new set representing the union of sets a and b; that is,
// those items which are members of either set a or set b.
func Union[T comparable](a, b Set[T]) Set[T] {
	out := make(Set[T])
	out.AddSet(a)
	out.AddSet(b)
	return out
}

// Intersection returns a new set representing the intersection of sets a and b; that is,
// those items which are members of both set a and set b.
func Intersection[T comparable](a, b Set[T]) Set[T] {
	out := make(Set[T])
	for v := range a {
		if _, ok := b[v]; ok {
			out[v] = struct{}{}
		}
	}
	return out
}

// Difference returns a new set representing the difference of sets a and b; that is,
// those items which are members of set a but not of set b.
func Difference[T comparable](a, b Set[T]) Set[T] {
	out := make(Set[T])
	for v := range a {
		if _, ok := b[v]; !ok {
			out[v] = struct{}{}
		}
	}
	return out
}

func (s *Set[T]) UnmarshalJSON(b []byte) error {
	var lst []T
	if err := json.Unmarshal(b, &lst); err != nil {
		return err
	}
	if *s == nil {
		*s = make(Set[T])
	}
	for _, v := range lst {
		(*s)[v] = struct{}{}
	}
	return nil
}

func (s Set[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Items())
}
