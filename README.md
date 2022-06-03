# set

[![Go Reference](https://pkg.go.dev/badge/github.com/jaz303/set.svg)](https://pkg.go.dev/github.com/jaz303/set)

This module provides a Go implementation of a generic set data type along with a small library of common operations.

Installation:

```
go get -u github.com/jaz303/set
```

## Usage Examples

### Create a new empty set

```go
s := make(set.Set[int])
```

### Create a set from known values

```go
s := set.Of("foo", "bar", "baz")
```

### Query the set

```go
s.Empty() // returns true if s is empty, false otherwise
s.Size() // return number of items in s
s.Contains(1) // return true if s contains item, false otherwise
s.ContainsSlice([]int{1,2,3}) // returns true if s contains all items in slice, false otherwise
s.Items() // returns a slice of all items in s
```

A set is implemented as a `map[T]struct{}` so the standard length/iteration operations are of course available:

```go
len(s) // return number of items in s

// iterate over s's contents
for v := range s {

}
```

### Add/remove items

```go
s.Add(1) // add a single item
s.AddSlice([]int{1,2,3}) // add all items from slice
s.AddSet(set.Of(4,5,6)) // add all items from other set

s.Remove(1) // remove a single item
s.RemoveSlice([]int{1,2,3}) // remove all items in slice
s.RemoveSet(set.Of(4,5,6)) // remove all items in other set

s.Clear() // remove all items from set
```

### Set operations

```go
set.Union(a, b) // returns a new set containing the union of sets a and b
set.Intersection(a, b) // returns a new set containing the intersection of sets a and b
set.Difference(a, b) // returns a new set containing the difference between sets a and b (i.e. a - b)
```

## JSON Support

The `Set` type implements the `json.Marshaler` and `json.Unmarshaler` interfaces; sets are encoded as JSON arrays.
