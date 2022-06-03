# set

Installation:

```
go get -u github.com/jaz303/set
```

## Examples

```go
func foo() {
    set := make(Set[int])
    set.Add(1)
    set.Add(2)
    set.AddSlice([]int{3, 4, 5})

    for val := range set {
        
    }
}
```