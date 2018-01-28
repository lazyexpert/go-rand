# go-rand
## Description
Very simple random int generator

## Dependencies
All you need is:
- build-essential (make)
- golang

## Build
Simply do:
```
make all
```

## Usage
There are two ways to generate random integer. 

### First
Simply pass maximum integer as a command line argument
```
./bin/rand 100
```

### Second
You pass two arguments, bounding the desired value.
Example: you want to generate number between 10 and 20.
Usage:
```
./bin/rand 10 20
```
