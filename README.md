## Pepelang!


[category]: <> (side-projects)
[date]: <> (2022/07/14)
[title]: <> (The pepelang programming language)
[color]: <> (green)

#### What is Pepelang?
 - Pepelang is an open source programming language supported by me 
 - Easy to learn and get started with (if you know spanish) 
 - Built-in sequential execution (get your concurrency out of here! who needs speed anyways?)
 - Open for PR and improvements!

### Project

#### Idea
Pepelang was born out of two personal desires, 
- Learning how to build my own programming language 
- More programming content available for the spanish speaking community

#### Concept
Pepelang is a programming language built and interpreted with/by Go.
It parses the source code in a Read-Eval-Print-Loop (REPL) environment and It has it's own token system, lexer, parser, AST and reuses the gargabe collection system from Go. 

Pepelang's code looks like this:

    var hey = fn() { 
	   retornar "hello world";
    }
    hey()
    
  Prints:
    "hello world"
    
## How to use?

#### Requirements
- Golang 

#### Install
Clone the repository into your local machine
`git clone https://github.com/Danielratmiroff/pepelang.git`
 
In the project's main folder, built it by running 
`go build -o .`

Run it with:
`./pepe` *or just execute the ".exe" if you're in windows*

**Linux**:
Add `pp` into your variable's path
*.bashrc or .zshrc*
`export PATH=$PATH:$HOME/PATH_TO_PROJECT/pepe`

### Features
You can create and run your own pepe programs using the ".pp" file extension. 

Create a `*.pp` file in any folder. `(e.g. hello.pp)`
Run them with:
`pepe hello.pp` or `./pepe hello.pp`


#### Syntax:
Pepelang supports the following:

Variables: Declared with the "var" keyword, can have either a string, number, boolean, array, dictionary or function value
`var hello = "world"`

Booleans: verdad (true) or falso (false)

Arithmetics supported: addition (+), substraction (-), multiplication (*), division (/), not operator (!), single operands (++) (--)
>> 4 + 2
6

>> !verdad
falso

>> var foo = 68
>> foo++
69

Arrays: Indexes are treated as expressions, thus, all of the following are valid:
>> var array = ["hello", 2, "world", 4]
>> array[3]
"world"

>> var array = [6 < 9 , 4 + 2]
>> expArray[1]
6

>> var array = [!verdad, !falso]
>> array[0]
falso

Dictionaries: Keys are literal values and values are parsed as expressions
>> var myDic = { "name": "daniel", "age": 69}; 
>> myDic["name"];
daniel 
>> myDic["age"]; 
69 

If/Else statements: Declared using the si (if) and sino (else), can take any expression value and evaluate the code accordingly
>> si ( 1 < 2 ) { retonar verdad }
verdad

>> si ( 42 == 69 ) { retonar verdad } sino { retornar falso };
falso

Return statements: Declared using the retonar keyword (as in previous example)

Function literals:
- Declared using the fn keyword
- Treated as high order functions (supports recursion and callback functions)
- Can take any number of parameters and such are treated as expressions.
- Support closure

>> var suma = fn(a, b) { a + b; };
>> suma(2, 2)
4

var sumaTotal = fn(f, x) { 
	 return sumaDos(sumaDos(x)); 
}; 
var sumaDos = fn(x) { 
	return x + 2; 
}; 

sumaTotal(sumaDos, 2); // => 6

### Built in functions (methods):
- Len: returns the length on a given value. Supported by strings, integers and arrays
>> len("Hello world")
12

>> var myArray = [1, 2, 3]
>> len(myArray)
2

- Puts: prints a value into the console
>> puts("print this message!")
print this message!

#### Array methods:
- First: gets the first element of a given array
- Last: gets the last element of a given array
>> var myArray = [1, 2, 3]
>> first(myArray)
1
>> last(myArray)
3

- Rest: removes the first element of a given array
>> var myArray = [1, 2, 3]
>> rest(myArray)
[2, 3]

- Push: inserts an element at the end of the array
>> push(myArray, "42")
[2, 3, "42"]

