package main

import "os"

func exit(x int) {
	os.Exit(x)
}

func testArgAssign(x int) int {
	x = 13
	return x
}

func testMinus() int {
	var x int = -1
	x = x * -5
	return x
}

func sum(x int, y int) int {
	return x + y
}

var globalint2 int
var globaluint8 uint8
var globaluint16 uint16
var globaluintptr uintptr

func testMisc() {
	var i13 int = 0
	i13 = testArgAssign(i13)
	var i5 int = testMinus()
	globalint2 = sum(1, i13 * i5)
	//var locali3 int
	var tmp int
	tmp = int(uint8('3' - '1'))
	tmp = tmp + int(globaluint16)
	tmp = tmp + int(globaluint8)
	tmp = tmp + int(globaluintptr)
	//locali3 = add1(tmp)
	exit(tmp)
}

func test() {
	testMisc()
}

func main() {
	test()
}
