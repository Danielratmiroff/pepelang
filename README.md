# Pepelang!
![version](https://img.shields.io/badge/version-0.3-green)
![coverage](https://img.shields.io/badge/coverage-85-yellowgreen)
![coolness](https://img.shields.io/badge/coolness-99-blue)


#### What is Pepelang?
 - Pepelang is an open source programming language supported by me 
 - Easy to learn and get started with *(if you know spanish)*
 - Built-in sequential execution *(get your concurrency out of here! who needs speed anyways?)*
 - Open for PRs and improvements!

## Project
Pepelang was born out of two personal desires, 
- Learning how to build my own programming language 
- More programming content available for the spanish speaking community

### Concept
Pepelang is a programming language built and interpreted with/by Go.

It parses the source code in a Read-Eval-Print-Loop (REPL) environment and It has it's own token system, lexer, parser, AST and reuses the gargabe collection system from Go. 

Pepelang's code looks like this:

```
var hola = fn() { 
   retornar "hello world";
}

hola()
```
    
*"hello world"*
    
## How to use?

### Requirements:
- [Golang](https://go.dev/) installed 

### Install
You'll need to [install Go](https://go.dev/doc/install)

```
git clone https://github.com/Danielratmiroff/pepelang.git
cd pepe
go install
```

### Usage
Call pepe in your terminal

```
$ pepe
```
If you want, you can also add an alias for this with echo "alias pp='pepe'" >> ~/.zshrc (or whichever rc file you're using).

**Please note: If you get an error claiming that lazygit cannot be found or is not defined, you may need to add ~/go/bin to your $PATH (MacOS/Linux), or %HOME%\go\bin (Windows). Not to be mistaken for C:\Go\bin (which is for Go's own binaries, not apps like pepe).**

## Features & Syntax
You can create and run your own pepe programs using the **".pp"** file extension. (Yes, ".pp" extension is already a thing, but I found it very funny so I decided to keep it!)

Create a `*.pp` file in any folder. `(e.g. hello.pp)`

Run it with:

`pp hello.pp # (if an alias was created)` or `./pepe hello.pp`

#### Pepelang supports the following functionality:

### Variables:

Declared with the `var` keyword, can have either a string, number, boolean, array, dictionary or function value

```
var hola = "mundo"
```

### Booleans:

```
verdad (true) or falso (false)
```

### Arithmetics supported:
- addition (+)
- substraction (-)
- multiplication (*)
- division (/)
- not operator (!)
- single operands (++) (--)
	
```
>> 4 + 2
6
```

```
>> !verdad
falso
```

```
>> var foo = 68
>> foo++
69
```

### Arrays:

Indexes are treated as expressions, thus, all of the following are valid:

```
>> var lista = ["hola", 2, "mundo", 4]
>> lista[3]
"mundo"
```

```
>> var lista = [6 < 9 , 4 + 2]
>> lista[1]
6
```
```
>> var lista = [!verdad, !falso]
>> lista[0]
falso
```

### Dictionaries:

Keys are literal values and values are parsed as expressions

```
>> var miDic = { "nombre": "daniel", "edad": 69}; 
>> miDic["nombre"];
daniel 
>> miDic["edad"]; 
69 
```

### If/Else statements:

Declared using the `si # (if)` and `sino # (else)`, can take any expression value and evaluate the code accordingly

```
>> si ( 1 < 2 ) { 
retonar verdad
}

verdad
```

```
>> si ( 42 == 69 ) { 
retonar verdad 
} sino { 
retornar falso
};

falso
```

### Return statements:

Declared using the `retonar` keyword (as in previous example)

### Function literals:
- Declared using the fn keyword
- Treated as high order functions (supports recursion and callback functions)
- Can take any number of parameters and such are treated as expressions.
- Support closure

```
>> var suma = fn(a, b) { a + b; };
>> suma(2, 2)
4
```

```
var sumaTotal = fn(f, x) { 
	 retornar sumaDos(sumaDos(x)); 
}; 
var sumaDos = fn(x) { 
	retornar x + 2; 
}; 

sumaTotal(sumaDos, 2); // => 6
```


### Built in functions:

- **Tam:**
Returns the length on a given value. Supported by strings, integers and arrays

```
>> tam("Hello world")
12
```

```
>> var lista = [1, 2, 3]
>> tam(lista)
2
```

- **Pon:**
Prints a value into the console

```
>> pon("Imprime este mensaje!")
imprime este mensaje!
```

- **Primero:**
Gets the first element of a given array

- **Ultimo:**
Gets the last element of a given array

```
>> var lista = [1, 2, 3]
>> primero(lista)
1
>> ultimo(lista)
3
```

- **Rest:**
Removes the first element of a given array

```
>> var lista = [1, 2, 3]
>> rest(lista)
[2, 3]
```

- **Empuja:**
Inserts an element at the end of the array

```
>> empuja(lista, "42")
[2, 3, "42"]
```

<br>

Enjoy!! 😄✨ (Spanish documentation coming soon! 
