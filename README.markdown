Interfaces bridge the gap between the fluidity of dynamically typed languages and the rigidity of statically typed languages. To demonstrate this, consider the differences between the duck.\* files in this repository. Each file is equivalent to the others but uses the different language features to accomplish the same task of making the object or struct quack.

Both Python and Javascript rely on a concept of "duck typing" which is common among dynamically typed languages. It harkens back to the adage:

> If it looks like a duck, swims like a duck, and quacks like a duck, then it probably is a duck.

Essentially, an argument passed into a function will work so long as all the methods called on that argument are valid. In duck.py and duck.js, the `Quack()` method is called on each object passed into `MakeQuack()`. The Duck quacks, the Cat quacks (don't ask), and the Dog doesn't. In Python, calling a non-existent method results in a run-time error and kills the program; in JS, the program continues but an undefined value is returned which can cause problems later on.

In a statically typed language, we cannot pass anything as a valid argument to a function. The data type of the argument must match the parameter. This would result in a lot of redundant code and make coding a nightmare. So, two concepts were devised to accommodate a degree of fluidity: interfaces and generics.

Go does not have generics; those are a milestone for Go 2.0 and an implementation is in work. According to the [proposal](https://go.googlesource.com/proposal/+/refs/heads/master/design/go2draft-type-parameters.md), Go's generics will be a variation of interfaces.

Interfaces define a [method set](https://golang.org/ref/spec#Method_sets). Any type (e.g. struct, map, array, int, etc.) implementing all of the methods of an interface can be used as that interface. In go_duck_simple/duck.go, the Quacker interface defines a single method - `Quack() string` - both Duck and Cat structs have a `Quack()` method matching the signature in Quacker. Therefore, both Duck and Cat can be passed into `MakeQuack()`which requires a Quacker argument.

[Method sets](https://golang.org/ref/spec#Method_sets) of types are more complex than they initially seem.

> The method set of any other type T consists of all methods declared with receiver type T. The method set of the corresponding pointer type *T is the set of all methods declared with receiver *T or T (that is, it also contains the method set of T).

The [receiver in a method declaration](https://golang.org/ref/spec#Method_declarations) determines whether a method is in the reference type method set (T) or in the point type method set (*T). A receiver can be declared in the following variations:

1. `func (t Type) Do()`
2. `func (Type) Do()`
3. `func (t *Type) Do()`
4. `func (*Type) Do()`

A reference type receiver (1) associates the function as a method of the type and (2) grants access to the fields of an instance of that type (so long as it has a variable declaration). A pointer type receiver can do everything a reference type receiver can do *and* it can change the instance's field values. In go_duck_nuanced/duck.go, a more developed implementation demonstrates the complexities.
