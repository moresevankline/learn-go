<!-- Compile and install the application -->
# Compile and install the application

Step 1: Do what is done in day_8, and in day_7 do only the `hello` directory

Step 2: Go to `hello` directory and compile the code into an executable

```bash
go build
```

Step 3: Confirm if the code works

```bash
./hello
```

Step 4: Discover the Go install path, where the go command will install the current package.

```bash
go list -f '{{.Target}}'
```

Step 5: Add the Go install directory to your system's shell path.

```bash
export PATH=$PATH:/path/to/your/install/directory
```

> The path is only up to `bin`, don't include the name

Step 6: Once updated the shell path, compile and install the package.

```bash
go install
```

Step 7: Run the application by typing its name

```bash
hello
```

## commands use

```bash
go build
go install
```
