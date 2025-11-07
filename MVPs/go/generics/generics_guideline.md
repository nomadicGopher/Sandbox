# Go Generics Learning Guide

This guide will help you learn how to use generics in Go, from the basics to advanced usage with interfaces.

---

## Table of Contents
1. [Introduction to Generics](#introduction-to-generics)
2. [Basic Syntax](#basic-syntax)
3. [Type Parameters](#type-parameters)
4. [Generic Functions](#generic-functions)
5. [Generic Types](#generic-types)
6. [Constraints](#constraints)
7. [Type Sets](#type-sets)
8. [Advanced: Generics and Interfaces](#advanced-generics-and-interfaces)
9. [Best Practices](#best-practices)
10. [Further Reading](#further-reading)

---

## Introduction to Generics
Generics allow you to write flexible and reusable code by parameterizing types. Introduced in Go 1.18, generics enable functions and types to operate on different data types without sacrificing type safety.

## Basic Syntax
A generic function or type uses square brackets `[]` to specify type parameters:

```go
func PrintSlice[T any](s []T) {
    for _, v := range s {
        fmt.Println(v)
    }
}
```

## Type Parameters
- `T` is a type parameter.
- `any` is a constraint (alias for `interface{}`) meaning any type is allowed.

## Generic Functions
You can define functions that work with any type:

```go
func Swap[T any](a, b T) (T, T) {
    return b, a
}
```

## Generic Types
You can define types (structs, etc.) with type parameters:

```go
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}
```

## Constraints
Constraints restrict the set of types that can be used as type parameters:

```go
func Add[T int | float64](a, b T) T {
    return a + b
}
```

## Type Sets
A type set is a list of types that satisfy a constraint:

```go
type Adder interface {
    type int, float64
}

func Add[T Adder](a, b T) T {
    return a + b
}
```

## Advanced: Generics and Interfaces
You can combine generics with interfaces for powerful abstractions.

### Example: Generic Interface Constraint
```go
type Stringer interface {
    String() string
}

func PrintStringer[T Stringer](s T) {
    fmt.Println(s.String())
}
```

### Example: Generic Data Structure with Interface Constraint
```go
type Equaler interface {
    Equal(other any) bool
}

type Set[T Equaler] struct {
    items []T
}

func (s *Set[T]) Add(item T) {
    for _, v := range s.items {
        if v.Equal(item) {
            return
        }
    }
    s.items = append(s.items, item)
}
```

## Best Practices
- Use generics for code reuse, not for everything.
- Keep constraints as simple as possible.
- Document your type parameters and constraints.

## Further Reading
- [Go Blog: Generics](https://go.dev/blog/intro-generics)
- [Go Generics Proposal](https://go.dev/design/43651-type-parameters)
- [Effective Go: Generics](https://go.dev/doc/effective_go#generics)