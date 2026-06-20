## `sort` a playlist of songs
Interface to comply to
```sh
type Interface interface {
  Len() int
  Less(i, j int) bool
  Swap(i, j int)
} 
```
To change Interface's `Less()` function
```sh
package sort
# struct wraps around an Interface => get Interface functions
type reverse struct { Interface } 
# override Less() function of old Interface
func (r reverse) Less(i, j int) bool { return r.Interface.Less(j,j) }
# function that convert old Interface to new object => implicitly new Interface (with new Less())
func Reverse(data Interface) Interface { return reverse(data) }

```sh
go run main.go
```
