package chess

import (
	"math/bits"
)

/*
 Square- & Piecewise representation of a chess board:
 Starting Square is a8, last is h1:
 ----------------------------------------------------
					FILE
		    a  b  c  d  e  f  g  h
		 .-------------------------.
	   8 | a8 b8 c8 d8 e8 f8 g8 h8 |
	   7 | a7 b7 c7 d7 e7 f7 g7 h7 |
	R  6 | a6 b6 c6 d6 e6 f6 g6 h6 |
	A  5 | a5 b5 c5 d5 e5 f5 g5 h5 |
	N  4 | a4 b4 c4 d4 e4 f4 g4 h4 |
	K  3 | a3 b3 c3 d3 e3 f3 g3 h3 |
	   2 | a2 b2 c2 d2 e2 f2 g2 h2 |
	   1 | a1 b1 c1 d1 e1 f1 g1 h1 |
	     '-------------------------'
 ------------------------------------------------------
 Bitboard is a 64bit int, having 1 bit for each square:
 relative to a piece-type we can save its positionings.
 ==> {a8}{b8}{c8}...{h8}{a7}{b7}...{h7}...{h1}
*/
type Bitboard uint64

func MakeBitboard(squares map[Square]bool) (b Bitboard) {
	for sq := range squares {
		b += (1 << sq)
	}
	return b
}

func (b Bitboard) IsSet(s Square) bool {
	return (bits.RotateLeft64(uint64(b), -int(s)) & 1) == 1
}

// helper function for testing
func (b Bitboard) areSet(sqs ...Square) bool {
	for _, sq := range sqs {
		if !b.IsSet(sq) {
			return false
		}
	}
	return true
}
