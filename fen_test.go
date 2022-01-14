package chess

import (
	"reflect"
	"testing"
)

// TODO: appearance / styling -> maybe subpackage for testcases and -variables?

type rankTest struct {
	name string
	r Rank
	rStr string
	err bool
	m map[File]Piece
}
func TestDecodeRank(t *testing.T) {
	tests := []rankTest{
		{"Full row", Rank8, "rnbqkbnr", false, map[File]Piece{
			FileA: BlackRook, FileB: BlackKnight, FileC: BlackBishop, FileD: BlackQueen,
			FileE: BlackKing, FileF: BlackBishop, FileG: BlackKnight, FileH: BlackRook,
		}},
		{"Mixed row", Rank7, "2P1rNBq", false, map[File]Piece{
			FileC: WhitePawn, FileE: BlackRook, FileF: WhiteKnight, FileG: WhiteBishop, FileH: BlackQueen,
		}},
		{"Empty row", Rank2, "8", false, map[File]Piece{}},
		{"Too short", Rank3, "3P", true, nil},
		{"Too long", Rank5, "5rrrr", true, nil},
		{"Invalid symbol", 15, "elmo", true, nil},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m, err := decodeRank(test.r, test.rStr)

			if (err != nil) != test.err {
				t.Fatalf("Invalid throwing of error %v when the appearance of an error should be %v", err, test.err)
			}
			if !reflect.DeepEqual(m, test.m) {
				t.Fatalf("Invalid mapping: %v instead of %v", m, test.m)
			}
		})
	}
}
// fen coded -> *board + error

type boardTest struct {
	name string
	bStr string
	err bool
	b *Board
}

var (
	petroffSquares = map[Square]Piece{
		H8: BlackRook, F8: BlackBishop, E8: BlackKing, D8: BlackQueen, 
		C8: BlackBishop, B8: BlackKnight, A8: BlackRook,
		H7: BlackPawn, G7: BlackPawn, F7: BlackPawn, C7: BlackPawn, 
		B7: BlackPawn, A7: BlackPawn,
		D6: BlackPawn,
		E4: BlackKnight, D4: WhitePawn,
		F3: WhiteKnight,
		H2: WhitePawn, G2: WhitePawn, F2: WhitePawn, 
		C2: WhitePawn, B2: WhitePawn, A2: WhitePawn,
		H1: WhiteRook, F1: WhiteBishop, E1: WhiteKing, D1: WhiteQueen, 
		C1: WhiteBishop, B1: WhiteKnight, A1: WhiteRook,
	}
	petroff = FromMap(petroffSquares)
)

func TestBoardFromFEN(t *testing.T) {
	tests := []boardTest{
		{"Turn Zero", "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR", false, &baseBoard},
		{"Petroff", "rnbqkb1r/ppp2ppp/3p4/8/3Pn3/5N2/PPP2PPP/RNBQKB1R", false, petroff},
		{"9 Ranks", "8/8/8/8/8/8/8/8/8", true, nil},
		{"7 Ranks", "8/8/8/8/8/8/8", true, nil},
		{"Invalid Rank", "rnbqkbnr/pppppppp/8/9/8/8/PPPPPPPP/RNBQKBNR", true, nil},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			bPtr, err := BoardFromFEN(test.bStr)

			if (err != nil) != test.err {
				t.Fatalf("Invalid throwing of error %v when the appearance of an error should be %v", err, test.err)
			}
			if !reflect.DeepEqual(bPtr, test.b) {
				t.Fatalf("Created wrong board %+v instead %+v", bPtr, test.b)
			}
		})
	}
}