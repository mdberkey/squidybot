package main

import (
	"fmt"
)

type Bitboard uint64
type Piece uint
type Side uint
type Square uint

const (
	WHITE Side = 0
	BLACK Side = 1
	BOTH Side = 2
	NUM_OF_PIECE_TYPES uint = 6
	NUM_OF_CASTLING_PERMISSIONS uint = 16
	NUM_OF_SQUARES uint = 64
	NUM_OF_FILES uint = 8
	NUM_OF_RANKS uint = 8
	CASTLING_WK uint8 = 1
	CASTLING_WQ uint8 = 2
	CASTLING_BK uint8 = 4
	CASTLING_BQ uint8 = 8
	CASTLING_ALL uint8 = 15
	EMPTY Bitboard = 0
	MAX_GAME_MOVES uint = 2048
	MAX_LEGAL_MOVES uint = 255
	MAX_PLY int8 = 125
	MAX_MOVE_RULE uint8 = 100
	FEN_START_POSITION string = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
)  

type board struct {
	bb_pieces [BOTH][NUM_OF_PIECE_TYPES]Bitboard
	bb_sides [BOTH]Bitboard
	game_state int // TODO: implement
	history int // TODO: implement
	piece_list [NUM_OF_SQUARES]Piece
}

func newBoard() {
	temp_pieces := board{}.bb_pieces
}

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
	printBitboard(123456)
	newBoard()
}  