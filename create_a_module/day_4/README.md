<!-- Call code from another module -->
# Call code from another module

Step 1: Do what is done in Day 3

Step 2: Create a directory. This is the caller.

```bash
mkdir hello
```

Step 3: Go to the created directory

```bash
cd hello
```

Step 4: Enable dependency tracking

```bash
go mod init example.com/hello
```

Step 5: Create `hello.go` file

```bash
touch hello.go
```

Step 6: Write code to call `Hello` function and print the return value

```go
package main

import (
  "fmt"

  "example.com/greetings"
)

func main() {
  // Get a greeting message and print it.
  message := greetings.Hello("Gladys")
  fmt.Println(message)
}
```

Step 7: Edit `example.com/hello` module to use local `example.com/greetings` module

```bash
go mod edit -replace example.com/greetings=../greetings
```

Step 8: Fix errors by synchronizing `example.com/hello` module dependencies

```bash
go mod tidy
```

Run the code

```bash
go run .
```
