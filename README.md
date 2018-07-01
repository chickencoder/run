<div align="right">
    [![Gitter chat](https://badges.gitter.im/gitterHQ/gitter.png)](https://gitter.im/gitterHQ/gitter)
</div>

# Run 

**Run** :runner: is a procedural scripting language intended for writing quick programs to automate tasks and to carry out general purpose computation. Run is designed to be obvious to read, write and run.

```javascript
import maths

entity Circle {
    radius
}

Circle func calcArea(self) {
    return maths.Pi * maths.pow self.radius, 2
}
```



## Basic Syntax

The (current) syntax of the Run language is heavily influenced by Go, Python and Javascript.

*Note: All code snippet currently use Javascript highlighting*

```javascript
print 'Hello World' 
```

The 'Hello World' program above showcases the `print` funtion from the Run Standard Library being called with a single parameter. By default, all types and literals within the Run language will have a default format to be printed with the `print` funtion.

Here are some things to bare in mind whilst writing **Run** code...

- Whitespace never matters
- Source programs are all UTF8
- Identifiers can be any valid UTF8 sequence



### Type System

Run is **dynamically typed** with a **duck type system** meaning that there is no type safety. Declaring variables therefore only comes in two flavours: immutable (constants) and mutable. Immutable variables are declared with the `set` keyword where as mutable variables are declared with the `let` keyword. Once a variable has been declared, it can be reassigned only if is mutable, with a simple `=` operator. 

```javascript
set pi = 3.141592654

let tau = 2 * pi
    tau = 6.283185307       # This is perfectly allowed

pi = 3.1                    # This most definitely is not
```

Despite the lack of types, Run only has 5 concrete data types. They are `numbers`, `strings`, `booleans`, `lists`, `maps`.

```javascript
# Numbers
let age = 10.12932

# Strings
let message = "Run is Fun"

# Lists
let languages = "Go", "Python", "Javascript"

# Maps
let languageToMessage = go: message, python: message, javascript: message
```

There will be more derived data types developed within the Standard Library such as Records.

`nil` is also a concrete type.

### Control Structures

One of the main control structures within the language is the `If Statement` which is has a familiar syntax and usage.

```javascript
let x = 10

if x < 20 {
    print "Run is fun"
} else {
    print "Run is not fun"    # This will (obviously) never run
}
```

It should also be noted at this point that funtions, when called with parameters will only consume the next token in the source unless followed by a comma. funtions can only be called without parameters using the `comma` syntax.

```javascript
print, # This will print a new line
```

The next notable control structure is the `for` statement which acts as both a conditional loop (often known as a while loop), and an iterative loop where by it can iterate over only `map` and `list` concrete types; with a special syntax. If no expression is provided to the `for` loop, just like in the Go Programming Language, the code will loop forever.

```Javascript
for {
   print "Forever" 
}
```

If the `for` loop is passed a comparison or truthy expression then it will only loop for as long as that condition is met.

```javascript
let name = getName,
for name is "Gregory" {
    print "Your name is still gregory"
}
```

Finally, if the `for` loop is passed an `of-iterator` expression then the `for` loop will iterate over a concrete collection type by assigning the mutable variable provided each time to the current item in the collection.

```javascript
for name of "Tom", "Dick", "Harry" {
    print concat "Hello", name
}
```



### Functions

funtions can be either **named** or **anonymous** and **pure** or **impure**. **Named** funtions are declared with a constant name where as **anonymous** funtions omit the name and act like values. A **pure** funtion is simply one that does not mutate any external state. funtions are always declared with the `fun` keyword.

```javascript
fun sayHello() {
    print "No."
}

fun pleaseSayHello() {
    print "Oh fine. Hello"
}
```

Technically speaking, the funtions above are both procedures since they do not take any arguments or return any values. To return values, the `return` keyword is used.



### Entities

**Run** does not support full object oriented programming, however, **Run** does provide an interface for describing *entities*.

```javascript
entity Animal {
    mass, species, favColour
}
```

Entities can also have methods. Strictly speaking, methods aren't tethered to entities like in most other languages. Instead, **Run** provides syntactic sugar whereby funtions can be defined that are automatically passed an instance of an entity. The instance can be reffered to using the `self` keyword. These special method-funtions can also only be called using the dot-syntax.

```javascript
Animal fun eat(f) {
    self.mass = inc(self.mass, 1)
}
```

Creating new instances of entities is the same as calling a function named after the entity name. All parameters passed must be in the order that they have been declared within the entity declaration.

```javascript
let dog = Animal(40, "lupus familiaris", "blue")
dog.eat()
```

### Protocols

*coming soon*

```javascript
protocol Shape {
    Area()
    Dimensions()
}

entity Square is Shape {
    length
}

Square fun Area(self) {
    return self.length * self.length
}
```



### Modules

**Run**'s module system is heavily inspired by the Python module system for simplicity sake. Every run source file (`*.run`) is considered a module. Modules are imported with the `import` keyword and only values exported with the `export` keyword can be accessed.

```javascript
# module.run
fun calculateArea(r) {
    return 3.14 * r * r
}

export calculateArea
```

```javascript
# main.run
import module
print module.calculateArea 10
```

Standard Library modules are also available through the import keyword. It should be noted that only local modules can be imported if they exist within the same directory, or a lower directory than that of the source file.



## Contributing

The **Run** language doesn't exist yet. It's merely a rumour. If you would like to suggest syntax changes, admendments or even features, get in contact. You can post an issue or PR if you're reading this on GitHub. Bare in mind that ideas will only be taken on board if the core developer aggrees.



Copyright (c) 2018 Jesse Sibley. All rights reserved.

This document, along with the entire source code for the 'Run' project is licensed under the BSD-3-Clause License. See `LICENSE` file for more details.
