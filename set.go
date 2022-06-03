package set

type Set[T comparable] map[T]struct{}

func Of[T comparable](vs ...T) Set[T] {
	out := make(Set[T])
	for _, v := range vs {
		out[v] = struct{}{}
	}
	return out
}

func (s Set[T]) Empty() bool       { return len(s) == 0 }
func (s Set[T]) Size() int         { return len(s) }
func (s Set[T]) Contains(v T) bool { _, ok := s[v]; return ok }

func (s Set[T]) ContainsSlice(vs []T) bool {
	for _, v := range vs {
		if _, ok := s[v]; !ok {
			return false
		}
	}
	return true
}

func (s Set[T]) Clear() {
	for k := range s {
		delete(s, k)
	}
}

func (s Set[T]) Add(v T) {
	s[v] = struct{}{}
}

func (s Set[T]) AddSlice(vs []T) {
	for _, v := range vs {
		s[v] = struct{}{}
	}
}

func (s Set[T]) AddSet(vs Set[T]) {
	for v := range vs {
		s[v] = struct{}{}
	}
}

func (s Set[T]) Remove(v T) {
	delete(s, v)
}

func (s Set[T]) RemoveSlice(vs []T) {
	for _, v := range vs {
		delete(s, v)
	}
}

func (s Set[T]) RemoveSet(vs Set[T]) {
	for v := range vs {
		delete(s, v)
	}
}

func (s Set[T]) Items() []T {
	out := make([]T, 0, len(s))
	for k := range s {
		out = append(out, k)
	}
	return out
}

func Union[T comparable](a, b Set[T]) Set[T] {
	out := make(Set[T])
	out.AddSet(a)
	out.AddSet(b)
	return out
}

func Intersection[T comparable](a, b Set[T]) Set[T] {
	out := make(Set[T])
	for v := range a {
		if _, ok := b[v]; ok {
			out[v] = struct{}{}
		}
	}
	return out
}

func Difference[T comparable](a, b Set[T]) Set[T] {
	out := make(Set[T])
	for v := range a {
		if _, ok := b[v]; !ok {
			out[v] = struct{}{}
		}
	}
	return out
}
