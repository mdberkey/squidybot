package main

import "fmt"


func printBitboard(bitboard uint64) {
	const lastBit uint64 = 63
	for rank := 0; rank < 8; rank++ {
		for file := 7; file >= 0; file-- {
			mask := uint64(1) << (lastBit - uint64(rank * 8) - uint64(file));
			char := '0'
			if bitboard & mask != 0 {
				char = '1'
			}
			fmt.Printf("%c ", char)
		}
		fmt.Println();
	}
}

func main() {
	printBitboard()
}  