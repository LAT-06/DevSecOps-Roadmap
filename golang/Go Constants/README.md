```golang
package main
import ("fmt")

const A = 1

func main() {
  fmt.Println(A)
}
```
> Note: In this case, the type of the constant is inferred from the value (means the compiler decides the type of the constant, based on the value).


When a constant is declared, it is not possible to change the value later:

```golang
package main
import ("fmt")

func main() {
  const A = 1
  A = 2
  fmt.Println(A)
}
```
Result:
```bash
./prog.go:8:7: cannot assign to A
```
Multiple Declaration

```golang
package main
import ("fmt")

const (
  A int = 1
  B = 3.14
  C = "Hi!"
)

func main() {
  fmt.Println(A)
  fmt.Println(B)
  fmt.Println(C)
}
```
