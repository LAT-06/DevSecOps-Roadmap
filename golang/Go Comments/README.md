For comments, it can be used `//` and `/* */`

```golang
// This is a comment
package main
import ("fmt")

func main() {
  fmt.Println("Hello World!")  // This is a comment
}
```

Or
```golang
package main
import ("fmt")

func main() {
  /* The code below will print Hello World
  to the screen */
  fmt.Println("Hello World!")
}
```

> Tip: It is up to you which you want to use. Normally, we use // for short comments, and /* */ for longer comments.