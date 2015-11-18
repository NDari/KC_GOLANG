## Should go be your next?

---

### A little about me

- Physics background, Started coding in FORTRAN, then Fortran and C
- Realized that I love programming with Perl
- Switched to Python for its scientific stack, and better object support
- switched to Go for the speeeeeed!

---

### What kind of programming language is Go?

- Strong, static types
- Enforced by a **very** fast compiler
- Garbage collected
- Concurrency supported at the language level, in the style of [Communicating Sequential Processes](http://www.usingcsp.com/cspbook.pdf)
- Functional programming facilities
- Object oriented programming facilities
- Very opinionated about... everything.

---

### Obligatory "hello world!"

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, KC Golang!\n")
}
```

[playground link](https://play.golang.org/p/cwOuObbyZ4)


- The "main" name of this package indicates that this will be an executable. Any other name will produce a static library.
- The importing is done by `import "package"` where the package name in in quotes.
- multiple packages can be imported by:

```go
import (
    "package1"
    "pakcage2"
)
```

- The entry point to any Go executable is the `main` function, just like C or C++.
- To get an exported item (Function, Struct, etc) out of a package, you use the dot notation
- The case of the first letter of an item in a package indicates if it is exported or not.
- All used packages must be imported, **and** all imported packages must be used. [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports)
- Automatic insertion of semi-colons, like Javascript. Nightmares, anyone? [gofmt](https://golang.org/cmd/gofmt/)


