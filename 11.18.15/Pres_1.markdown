# Should go be your next?

Kansas City Go-lang meetup

11/18/2015

---

## A little about me

- Physics background, Started coding in FORTRAN, then Fortran and C
- Realized that I love programming with Perl
- Switched to Python for its scientific stack, and better object support
- switched to Go for the speeeeeed!

---

## What kind of programming language is Go?

- Strong, static types
- Enforced by a **very** fast compiler
- Garbage collected
- Concurrency supported at the language level, in the style of [Communicating Sequential Processes](http://www.usingcsp.com/cspbook.pdf)
- Functional programming facilities
- Object oriented programming facilities
- Very opinionated about... everything.
- Focused on simplicity, and being as lightweight as possible.
- This pays off in the tooling.

---

## Obligatory "hello world!"

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, KC Golang!\n")
}
```

- The "main" name of this package indicates that this will be an executable. Any other name will produce a static library.
- The importing is done by `import "package"` where the package name in in quotes.
- The entry point to any Go executable is the `main` function, just like C or C++.
- To get an exported item (Function, Struct, etc) out of a package, you use the dot notation
- The case of the first letter of an item in a package indicates if it is exported or not.
- All used packages must be imported, **and** all imported packages must be used. [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports)
- Automatic insertion of semi-colons, like Javascript. Nightmares, anyone? [gofmt](https://golang.org/cmd/gofmt/)

---

## Functions and functional programming

- Functions are first class objects.

```go
package main

import "fmt"

func main() {
	square := func(i float64) float64 {
		return i * i
	}

	v := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	fmt.Println(mymap(square, v))

	odd := func(i int) bool {
		return i%2 == 0
	}

	t := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(myfilter(odd, t))
}

func mymap(f func(float64) float64, v []float64) []float64 {
	for i := 0; i < len(v); i++ {
		v[i] = f(v[i])
	}
	return v
}

func myfilter(f func(int) bool, v []int) []int {
	var res []int
	for i := 0; i < len(v); i++ {
		if f(v[i]) {
			res = append(res, v[i])
		}
	}
	return res
}
```

[playground link](https://play.golang.org/p/rnE7aYRGZl)

- Functions can have closures

```go
package main

import "fmt"

func main() {
	inc, sum := accumulator(), accumulator()
	for i := 0; i < 5; i++ {
		fmt.Println("inc", inc(1), "\nsum", sum(i))
	}
}

func accumulator() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
```

[playground link](https://play.golang.org/p/RahnonLxAX)

- Function with variable number of arguments (veradic functions) are allowed.

```go
package main

import "fmt"

func sum(nums ...int) {
    fmt.Print(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}

func main() {

    // Variadic functions can be called in the usual way
    // with individual arguments.
    sum(1, 2)
    sum(1, 2, 3)

    // If you already have multiple args in a slice,
    // apply them to a variadic function using
    // `func(slice...)` like this.
    nums := []int{1, 2, 3, 4}
    sum(nums...)
}
```

- All arguments are pass by value, except slices.
- Can pass by reference if you pass pointers to functions (more on this soon!)

---

## Structs

Structs are simply a way to group data together. They can be thought of as classes, in other languages with some caveats.

```go
package main

import "fmt"

type Person struct {
    Age  int
    Name string
}

type Employee struct {
    Person
    ID     int
}

func main() {
    p := Employee{Person{"John", 35}, 223341}
    fmt.Println(p)
}
```

[playground link](https://play.golang.org/p/-HfpOcNYpu)

- No inheritance, only composition.
- Person is embedded into Employee. Allows direct access to Person's fields from Employee (Employee.Age, etc)
- Also allows direct access to Person's methods as well (more on methods shortly)
- Can overwrite the fields from Person by implementing it in Employee. `Employee.Name != Employee.Person.Name`
- This is awful... [lets fix it](https://play.golang.org/p/PKPcmbOg8J)
- Cannot initialize fields. They are automatically "zero-ed" out.
- Get around this by hiding the default constructor, and rolling your own.

```go
package main

import "fmt"

type Person struct {
    Age  int
    Name string
}

type employee struct {
    Person
    ID     int
}

func NewImployee(id int) Employee {
    return employee{Person{}, ID: id}
}

func main() {
    p := NewEmployee(223341}
    fmt.Println(p)
}
```


---

## Methods

Methods in go are just syntactic sugar for functions. Consider the two functions below:

```go
func (p Person) SayHi() {
    fmt.Println("Hi! My name is", p.Name)
}

func SayHi2(p Person) {
    fmt.Println("Hi! My name is", p.Name)
}
```

```go
func (e Employee) ChangeID(newID int) {
    e.ID = newID
}
```

- The above function does **not** modify the employee. since everything is pass by value, a copy is made.
- We can [fix this too](https://play.golang.org/p/WBCg6A9bui)
- Methods can only be defined in the same package that defines the struct :(

---

## Interfaces

Interfaces are a collection of methods implemented by a Struct. This is the way go implements duck typing.


```go
type bro struct {
    name string
}

func (b bro) sup() {
    fmt.Println("sup brah! I am", b.name)
}
G
type dude struct {
    name string
}

func (d dude) sup() {
    fmt.Println("Hey Dude! I am", d.name)
}

type greeter interface {
    sup()
}

func greet(g greeter) {
    g.sup()
}

func main() {
    jim := bro{"jim"}
    joe := dude{"joe"}
    greet(jim)
    greet(joe)
}
```

[playground link](https://play.golang.org/p/gQLbGB3jHT)

- Interfaces are implicit: if you can do _this_, then you can be used _here_.
- Can build interfaces that take third-party libraries without modifying them. The compiler still checks for you.
- The empty interface is a catch all for everything... 

## Channels and Goroutines

- Each is awesome on its own, but together, they make Golang as powerful as it is.
- Dont communicate through sharing memory, but share memory through communicating.
- Channels are how processes can communicate between each other.
- goroutines are "very light weight threads", that can run concurrently.

[a silly example](https://play.golang.org/p/0GXa2IlI9c)

## Tools

- go vet
- go doc: generate documentation from code. [Example](https://godoc.org/github.com/stretchr/testify/assert)
- go get: light weight go package manager... of sorts). A good set up for Go:

```bash
export GOROOT=/usr/local/go
export GOPATH=$HOME/golib:$HOME/go
export GOBIN=$HOME/go/bin
export PATH=$GOBIN:$PATH
```

- go test: testing and benchmarking. [Example](https://github.com/NDari/numgo)
- go oracle: Code comprehension tool. [link](https://docs.google.com/document/d/1SLk36YRjjMgKqe490mSRzOPYEDe0Y_WQNRv-EiFYUyw/view)

## Resources

- [Golang tour](https://tour.golang.org)
- [Go by example](https://gobyexample.com/)
- [Concurrency is not Parallelism](https://www.youtube.com/watch?v=cN_DpYBzKso)
- [Advanced Concurrency Patterns](https://www.youtube.com/watch?v=QDDwwePbDtw)


