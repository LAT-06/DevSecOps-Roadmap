In Go, there are different types of variables, for example:

- `int`- stores integers (whole numbers), such as 123 or -123
- `float32`- stores floating point numbers, with decimals, such as 19.99 or -19.99
- `string` - stores text, such as "Hello World". String values are surrounded by double quotes
- `bool`- stores values with two states: true or false

## Declaring aka Creating Variables

Using `var` keyword
```golang
var <variablename> <type> = <value>
```
> Note: You always have to specify either `type` or `value` (or both).

Using `:=` sign

```golang
variablename := value
```

> Note: In this case, the type of the variable is inferred from the value (means that the compiler decides the type of the variable, based on the value).

> It is not possible to declare a variable using `:=`, without assigning a value to it.

Difference between `var` and `:=`

| var | := |
| --- | --- |
| Can be used inside and outside of functions | Can only be used inside functions |
| Variable declaration and value assignment can be done separately | Variable declaration and value assignment cannot be done separately (must be done in the same line) |

It can be declared multi variables by using `,`

