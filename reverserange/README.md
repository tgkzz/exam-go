## reverserange

### Instructions

Write a program that takes two numbers as arguments and prints the consecutive values that begins at the second number and end at the first number (including the values of those numbers !).

Errors should be handled.

If the number of arguments is different from 2 the program prints nothing.

### Usage :

```console
$ go run . 1 3
3 2 1
$ go run . -1 2 | cat -e
2 1 0 -1$
$ go run . 0 0
0
$ go run . 0 -3
-3 -2 -1 0
$ go run . 0 nan | cat -e
strconv.Atoi: parsing "nan": invalid syntax$
$
```