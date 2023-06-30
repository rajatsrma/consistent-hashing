package main

import (
	_ "fmt"
)

var alphabetMap = map[rune]int{
	rune('a'): 1,
	rune('b'): 2,
	rune('c'): 3,
	rune('d'): 4,
	rune('e'): 5,
	rune('f'): 6,
	rune('g'): 7,
	rune('h'): 8,
	rune('i'): 9,
	rune('j'): 10,
	rune('k'): 11,
	rune('l'): 12,
	rune('m'): 13,
	rune('n'): 14,
	rune('o'): 15,
	rune('p'): 16,
	rune('q'): 17,
	rune('r'): 18,
	rune('s'): 19,
	rune('t'): 20,
	rune('u'): 21,
	rune('v'): 22,
	rune('w'): 23,
	rune('x'): 24,
	rune('y'): 25,
	rune('z'): 26,
	rune('0'): 27,
	rune('1'): 28,
	rune('2'): 29,
	rune('3'): 30,
	rune('4'): 31,
	rune('5'): 32,
	rune('6'): 33,
	rune('7'): 34,
	rune('8'): 35,
	rune('9'): 36,
	rune('-'): 37,
}

func getHash(context string) int {
	var hashSum = 0
	for i := 0; i < len(context); i++ {
		hashSum += (alphabetMap[rune(context[i])] + hashSum)
	}

	return hashSum % 360
}
