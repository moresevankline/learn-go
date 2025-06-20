<!-- Get started with Go -->
# Get started with Go

Step 1: Create directory

```bash
mkdir hello
```

Step 2: Go to the directory

```bash
cd hello
```

Step 3: Enable dependency tracking

```bash
go mod init example/hello
```

Step 4: Create a go file

```bash
touch hello.go
```

Step 5: Write some hello world code in `hello.go`

```go
package main

import "fmt"

func main() {
  fmt.Println("Hello, World!")
}
```

Run the code

```bash
go run .
```

<!-- Additional -->
Step 6: Import external package `"rsc.io/quote"` in `hello.go`

```go
package main

import "fmt"

import "rsc.io/quote"

func main() {
  fmt.Println(quote.Go())
}
```

Step 7: Read error `could not import rsc.io/quote (no required module provides package "rsc.io/quote")compilerBrokenImport`

Step 8: Fix error by adding module requirements and sums

```bash
go mod tidy
```

Run the code

```bash
go run .
```
