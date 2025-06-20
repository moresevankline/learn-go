<!-- Create a module -->
# Create a module

Step 1: Create directory

```bash
mkdir greetings
```

Step 2: Go into the directory

```bash
cd greetings
```

Step 3: Enable dependency tracking

```bash
go mod init example.com/greetings
```

Step 4: Create `greetings.go` file

```bash
touch greetings.go
```

Step 5: Write a `Hello` function, `name` as parameter with `string` parameter type, and return `string`

```go
package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
  // Return a greeting that embeds the name in a message.
  message := fmt.Sprintf("Hi, %v. Welcome!", name)
  return message
}
```

## Function format in Go

```go
func {functionName}({parameter} {parameterType}) {returnType} {
  // code
}
```
