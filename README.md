# Optional Package for Go

The `optional` package provides a generic `Optional` type for Go, which represents a value that may or may not be present. This can be useful for handling optional values in a type-safe way.

## Features

- **Generic Type**: The `Optional` type can be used with any data type.
- **Presence Check**: Check if a value is present.
- **Value Retrieval**: Safely retrieve the value or handle the absence of a value.
- **Default Values**: Provide default values when the actual value is absent.

## Installation

To install the package, run:

```sh
go get example.com/optional
```

## Usage

### Importing the Package

```go
import "example.com/optional"
```

### Creating an Optional Value

```go
package main

import (
    "fmt"
    "example.com/optional"
)

func main() {
    // Create an Optional with a value
    optInt := optional.New(42)

    if optInt.IsPresent() {
        value, _ := optInt.Get()
        fmt.Println("Value:", value)
    } else {
        fmt.Println("No value present")
    }

    // Create an empty Optional
    emptyOpt := optional.Empty[int]()
    if emptyOpt.IsPresent() {
        value, _ := emptyOpt.Get()
        fmt.Println("Value:", value)
    } else {
        fmt.Println("No value present")
    }

    // Using OrElse to provide a default value
    defaultValue := emptyOpt.OrElse(100)
    fmt.Println("Default Value:", defaultValue)
}
```

### Example with a String

```go
package main

import (
    "fmt"
    "example.com/optional"
)

func main() {
    optStr := optional.New("Hello, World!")
    if optStr.IsPresent() {
        value, _ := optStr.Get()
        fmt.Println("Value:", value)
    } else {
        fmt.Println("No value present")
    }
}
```

## API

### `Optional` Struct

#### Methods

- **`New`**: Creates a new `Optional` containing the given value.

```go
  func New[T any](value T) Optional[T]
  ```

- **`Empty`**: Creates a new empty `Optional`.

```go
  func Empty[T any]() Optional[T]
  ```

- **`IsPresent`**: Returns true if the `Optional` contains a value.

```go
  func (opt Optional[T]) IsPresent() bool
  ```

- **`Get`**: Returns the value if present, or an error if not.

```go
  func (opt Optional[T]) Get() (T, error)
  ```

- **`OrElse`**: Returns the value if present, or the provided default value if not.

```go
  func (opt Optional[T]) OrElse(defaultValue T) T
  ```

## Testing

To run the tests, navigate to the directory containing your `go.mod` file and run:

```sh
go test ./...
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
