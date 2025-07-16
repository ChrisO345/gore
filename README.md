# gore

`gore` is a simple Go package that provides a set of testing utilities that integrate with the Go testing framework.

---

## Features

- **Simple Assertions** - Straightforward helpers for common test conditions.
- **Fluent Interface** - Write readable, chainable assertions.
- **Fail-Fast Option** - Includes a stricter `vital` mode for critical assertions.

---

## Installation

`gore` is available on GitHub and can be installed using Go modules:

```bash
go get github.com/chriso345/gore
```

## Usage

### Basic Assertions

Import and use the `assert` package to perform standard test assertions:

```go
import "github.com/chriso345/gore/assert"

func TestExample(t *testing.T) {
    result := 1 + 1
    assert.Equal(t, result, 2)
}
```

### Fluent Interface

Use a fluent API for more expressive and concise assertions:

```go
import "github.com/chriso345/gore/assert"

func TestExample(t *testing.T) {
    a := assert.That(t)
    a.Equal(1+1, 2)
}
```

### Vital Mode (Fail-Fast)

Use the `vital` package for assertions that halt execution immediately on failure (`t.Fatalf`):

```go
import "github.com/chriso345/gore/vital"

func TestCritical(t *testing.T) {
    vital.Equal(t, 1+1, 2)
}
```

Fluent Assertions are also available in `vital`:

```go
import "github.com/chriso345/gore/vital"

func TestCriticalFluent(t *testing.T) {
    v := vital.That(t)
    v.Equal(1+1, 2)
}
```

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
