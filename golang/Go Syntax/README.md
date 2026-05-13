## GO Syntax
A Go file consists of the following parts:
- Package declaration
- Import packages
- Functions
- Statements and expressions

```golang
package main
import ("fmt")

func main() {
  fmt.Println("Hello World!")
}
```
### Explanatioin
Line 1: Every Golang program is part of package and define this using the `package` keyword. The program belongs to main package.

Line 2: Import package (ex: fmt) `import("fmt")`

Line 3: `func main(){}` is a function.

Line 4: `fmt.Println()` is a function made available from the `fmt` package. It is used to output/print text

> Note: In Go, any executable code belongs to the `main` package.

## Go Statement
In golang the statement can be separated by hitting Enter or by using `;`

The left curly bracket `{` cannot come at the start of a line.

You cannot run this
```golang
package main
import ("fmt")

func main()
{
  fmt.Println("Hello World!")
}
```
