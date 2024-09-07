# Go programming

Having a single module (in the project) root directory is the **standard approach.**
Having multiple modules in a project is generally not recommended (unless perhaps a very huge project)

Review of the review.

## Single Module in Project Root Dir

Sample of Go project structure with a single go.mod
file in the root directory:

```txt
myproject/
├── go.mod           # Main module file
├── go.sum           # Module checksums (generated automatically)
├── main.go          # Main application entry point
├── pkg/
│   ├── greetings/   # Package for greeting messages
│   │   └── greetings.go
│   └── utils/       # Package for utility functions
│       └── utils.go
└── cmd/ 
    └── myapp/
        └── main.go   # Entry point for a command-line tool
```

go.mod:

```go
module myproject

go 1.21  // The Go version

require (
    // List any external dependencies here
)
```

greetings.go (pkg/greetings/greetings.go):

```go
package greetings // Package name matches directory

// Hello returns a greeting message.
func Hello(name string) string {
    return fmt.Sprintf("Hello, %s!", name)
}
```

utils.go (pkg/utils/utils.go):

```go
package utils

// IsEven checks if a number is even.
func IsEven(num int) bool {
    return num%2 == 0
}
```

main.go (main.go):

```go
package main

import (
    "fmt"
    "myproject/pkg/greetings"
    "myproject/pkg/utils"
)

func main() {
    message := greetings.Hello("Alice")
    fmt.Println(message)

    num := 10
    if utils.IsEven(num) {
        fmt.Println(num, "is even")
    }
}
```

myapp/main.go (cmd/myapp/main.go):

```go
package main // Another main package for a different executable

import (
    "fmt"
    "myproject/pkg/greetings"
)

func main() {
    message := greetings.Hello("CLI User")
    fmt.Println(message)
}
```

Explanation:

- go.mod: This file defines the module name (myproject) and tracks the project's dependencies.
- pkg: This directory contains internal packages that can be reused in the project.
  - greetings: Provides a function for creating greeting messages.
  - utils: Contains utility functions (in this example, a function to check if a number is even).
- cmd: This directory typically holds packages that define command-line tools.
  - myapp: A separate executable tool that uses the greetings package.
- main.go: Main entry point for the primary application. It uses the greetings and utils packages.

Key Points:

- Single Module: Entire project is managed as a single Go module with one go.mod file in root directory.
- Internal Packages: The pkg directory holds reusable packages that are imported using the module path (myproject/pkg/greetings, myproject/pkg/utils).
- Multiple Executables: We can have multiple executables within the project, each defined in a separate package main directory (like cmd/myapp in this example).

This is a flexible structure that can be scaled to accommodate larger projects.
