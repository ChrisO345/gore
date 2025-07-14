# gore

`gore` is a simple Go package that provides a set of testing utilities that integrate with the Go testing framework.

---

## Features

- **Test Assertions**: Provides a set of assertion functions to validate conditions in tests.
- **Fluent Interface**: Allows for a more readable and fluent way to write assertions.

---

## Installation

`gore` is available on GitHub and can be installed using Go modules:

```bash
go get github.com/chriso345/gore
```

## Usage

`gore/assert` provides a simple set of assertion functions that can be used in tests.

```go
func TestExample(t *testing.T) {
  result := 1 + 1
  assert.Equal(t, result, 2)
}
```

`gore/assert` also provides a way to create a more fluent interface for assertions:

```go
func TestExample(t *testing.T) {
  a := assert.That(t)
  result := 1 + 1
  a.Equal(result, 2)
}
```

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
