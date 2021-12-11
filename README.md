
Go
- Developed by Google
- It is complied language (simply have to compile that will generate machine code for you) similar to C/C++

Setup
- complier
- IDE (Integrated Development Environment)

Commands
```
go version
go mod init example/hello
go run .
```
Or
```
go run first.go
```

If you perform an action, you have to call a function or method in oops.

### Packages
Eveything in GO will required packages, use packages with name
eg: print, connect with internet

Common packages name: `fmt` (format), `io`, `sql`

### Code
- Excution is start with main function, create a function 
```
package main

import 'fmt'
func main(){
  fmt.Println('something')
}
```

### Variable
- Where we can change the value of data
- permanent data is stored in the database
- temporary data is stored in variables

Variable Features
1. Name of the variable
2. type of the variable
3. data stored inside the variable

```
var name_of_variable type_of_variable = value_assigned 
```
type_of_variable are:
- bool
- string
- int int8 int16 int32 int64
- uint uint8 uint16 uint32 uint32 uintptr
- byte //alias for uint8
- rune //alias for int32
- float32 float64

### Constant
```
const something = 'constant'
```

### Loops
```
i := 1
for i<=5 {
  fmt.Println("something", i)
  i++
}
```

### Global Scope
- functional level (local scope)
- package level (global scope)

### Exported Names

### Conditional
- if else
- switch 
