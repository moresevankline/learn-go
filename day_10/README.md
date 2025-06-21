<!-- Getting started with multi-module workspaces -->
# Getting started with multi-module workspaces

## Create a module for code

Step 1: Create a directory called `workspace`

```bash
mkdir workspace
cd workspace
```

Step 2: Create `hello` module

```bash
mkdir hello
cd hello
go mod init example.com/hello
```

Step 3: Add a dependency on the `golang.org/x/example/hello/reverse` package using `go get`

```bash
go get golang.org/x/example/hello/reverse
```

Step 4: Create `hello.go` file

```bash
touch hello.go
```

Step 5: Write code

```go
package main

import (
  "fmt"

  "golang.org/x/example/hello/reverse"
)

func main() {
  fmt.Println(reverse.String("Hello"), reverse.Int(24601))
}
```

Step 6: Run `hello` program

```bash
go run .
```

## Create the workspace

Step 7: Go to `workspace` directory

```bash
cd ..
```

Step 8: Initialize the `workspace`

```bash
go work init ./hello
```

Step 9: Run the program in the `workspace` directory

```bash
go run ./hello
```

## Download and modify the `golang.org/x/example/hello` module

Step 10: Clone the repository

```bash
git clone https://go.googlesource.com/example
```

Step 11: Add the module to the workspace

```bash
go work use ./example/hello
```

Step 11: Go to `workspace/example/hello/reverse` directory

```bash
cd example/hello/reverse
```

Step 12: Create `int.go` file

```bash
touch int.go
```

Step 13: Write code

```go
package reverse

import "strconv"

// Int returns the decimal reversal of the integer i.
func Int(i int) int {
  i, _ = strconv.Atoi(String(strconv.Itoa(i)))
  return i
}

```

Step 14: Go to `workspace/hello` directory

```bash
cd ../../../hello
```

Step 15: Modify `hello.go` file

```go
package main

import (
  "fmt"

  "golang.org/x/example/hello/reverse"
)

func main() {
  fmt.Println(reverse.String("Hello"), reverse.Int(24601))
}
```

Step 16: Go to `workspace` directory

```bash
cd ..
```

Step 17: Run the code

```bash
go run ./hello
```

### multi-module workspace

Writing code in multiple modules at the same time and easily build and run code in those modules

### commands

```bash
go get <package> # add a dependency package
go work init <directory> # tells Go to create a go.work file containing the modules in the directory
go work use <directory> # adds a new module to the go.work file
```
