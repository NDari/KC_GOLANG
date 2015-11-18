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

---

## Obligatory "hello world!"

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
- The entry point to any Go executable is the `main` function, just like C or C++.
- To get an exported item (Function, Struct, etc) out of a package, you use the dot notation
- The case of the first letter of an item in a package indicates if it is exported or not.
- All used packages must be imported, **and** all imported packages must be used. [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports)
- Automatic insertion of semi-colons, like Javascript. Nightmares, anyone? [gofmt](https://golang.org/cmd/gofmt/)

---

## Functions and functional programming

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

- Functions are first class objects.
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

func (e Employee) ChangeID(newID int) {
    e.ID = newID
}
```

- The above function does **not** modify the employee. since everything is pass by value, a copy is made.
- We can [fix this too](https://play.golang.org/p/WBCg6A9bui)

---

## resources

- [Go by example](https://gobyexample.com/)
- [Golang tour](https://tour.golang.org)

