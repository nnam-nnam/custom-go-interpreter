# Monkey Language Go Interpreter

## Background
This repository hosts the code from Thorsten Ball's book [
Writing An Interpreter In Go
](https://interpreterbook.com/). The goal is to use this project as an introduction to Golang and additionally to learn more about the design of programming languages.

## Language Features
Currently the Monkey language allows us to bind expressions to identifiers using `let` statements:
```
let name = "gorilla";

let array = [1, 2, 4, 8, 16];

let isAwesome = true;

let longExpression = (5 * 10 + 3) / 2 - 5;

let capitalCities = {"France" : "Paris"};
```

There's also support for functions:
```sh
>> let addTwo = fn(x) { x + 2 };
>> addTwo(3)
5
```

Monkey also has built-in array and hash data types:
```sh
>> let a = [1, 2, 3, 4];
>> a
[1, 2, 3, 4]
>> first(a)
1
>> last(a)
4
>> rest(a)
[2, 3, 4]
>> push(a, 5)
[1, 2, 3, 4, 5]
>> let h = {"one" : 1};
>> h["one"]
1
>> a[0]
1
```

## Installation
If you would like to use the REPL to run Monkey commands, you can clone this repo and run the program, assuming your machine has a Go version of >=1.20.

Just clone and run this at the root:
```sh
cd src/monkey && go run main.go
```

Alternatively you can run this project in a Docker container if that's more suitable:
```sh
docker build -t monkey .
docker run -it --rm monkey
```
