<!-- Add a test -->
# Add a test

Step 1: Do what is done in Day 7 but only the `greetings` directory

Step 2: Create `greetings_test.go` file inside `greetings` directory

```bash
touch greetings_test.go
```

Step 3: Write test functions for `greetings.go`

```go
package greetings

import (
  "regexp"
  "testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
  name := "Gladys"
  want := regexp.MustCompile(`\b` + name + `\b`)
  msg, err := Hello("Gladys")
  if !want.MatchString(msg) || err != nil {
    t.Errorf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
  }
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
  msg, err := Hello("")
  if msg != "" || err == nil {
    t.Errorf(`Hello("") = %q, %v, want "", error`, msg, err)
  }
  }
```

Step 4: Run the test command to execute test

```bash
go test
go test -v
```

Step 5: Modify `greetings.go` to make the test fail

```go
package greetings

import (
  "errors"
  "fmt"
  "math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
  // If no name was given, return an error with a message.
  if name == "" {
    return "", errors.New("empty name")
  }

  // If a name was received, return a value that embeds the name
  // in a greeting message.
  // message := fmt.Sprintf(randomFormat(), name)
  message := fmt.Sprint(randomFormat())
  return message, nil
}

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
  // A map to associate names with messages.
  messages := make(map[string]string)
  // Loop through the received slice of names, calling
  // the Hello function to get a message for each name.
  for _, name := range names {
    message, err := Hello(name)
    if err != nil {
      return nil, err
    }
    // In the map, associate the retrieved message with
    // the name.
    messages[name] = message
  }
  return messages, nil
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
  // A slice of message formats.
  formats := []string{
    "Hi, %v. Welcome!",
    "Great to see you, %v",
    "Hail, %v! Well met!",
  }

  // Return a randomly selected message format by specifying
  // a random index for the slice of formats.
  return formats[rand.Intn(len(formats))]
  }
```

Step 6: Run the test command to execute test

```bash
go test
go test -v
```

new module

```go
import testing /* for testing */
```

> Ending a file's name with `_test.go` tells the go test command that this file contains test functions.
