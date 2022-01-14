package chess

import (
	"reflect"
	"testing"
)

var (
	basePosition = map[Square]Piece{
		H8: BlackRook, G8: BlackKnight, F8: BlackBishop, E8: BlackKing, 
		D8: BlackQueen, C8: BlackBishop, B8: BlackKnight, A8: BlackRook,
		H7: BlackPawn, G7: BlackPawn, F7: BlackPawn, E7: BlackPawn,
		D7: BlackPawn, C7: BlackPawn, B7: BlackPawn, A7: BlackPawn,
		H2: WhitePawn, G2: WhitePawn, F2: WhitePawn, E2: WhitePawn, 
		D2: WhitePawn, C2: WhitePawn, B2: WhitePawn, A2: WhitePawn,
		H1: WhiteRook, G1: WhiteKnight, F1: WhiteBishop, E1: WhiteKing,
		D1: WhiteQueen, C1: WhiteBishop, B1: WhiteKnight, A1: WhiteRook,
	}
	baseWPawns = Bitboard(255<<48)
	baseBPawns = Bitboard(255<<8)
	baseWKnights = Bitboard(66<<56)
	baseBKnights = Bitboard(66<<0)
	baseWBishops = Bitboard(36<<56)
	baseBBishops = Bitboard(36<<0)
	baseWRooks = Bitboard(129<<56)
	baseBRooks = Bitboard(129<<0)
	baseWQueens = Bitboard(8<<56)
	baseBQueens = Bitboard(8<<0)
	baseWKing = Bitboard(16<<56)
	baseBKing = Bitboard(16<<0)
	baseWhite = Bitboard(65535<<48)
	baseBlack = Bitboard(65535<<0)
	baseTotal = baseWhite+baseBlack
	empty = Bitboard(4_294_967_295<<16)
	baseBoard = Board{
		wPawns: baseWPawns, bPawns: baseBPawns,
		wKnights: baseWKnights,bKnights: baseBKnights,
		wBishops: baseWBishops,bBishops: baseBBishops,
		wRooks: baseWRooks,bRooks: baseBRooks,
		wQueens: baseWQueens,bQueens: baseBQueens,
		wKing: baseWKing,bKing: baseBKing,
		wPieces: baseWhite, bPieces: baseBlack,
		pieces: baseTotal, empty: empty,
	}
)

func TestFromMap(t *testing.T) {
	board := *FromMap(basePosition)

	if all := baseTotal + empty + 1; all != 0 {
		t.Fatalf("Learn Math: %d instead of 0", all)
	}

	if !reflect.DeepEqual(board, baseBoard) {
		t.Fatalf("Created wrong board: %+v instead of %+v", board, baseBoard)
	}
}

type pieceBitboard struct {
	name string
	p Piece
	bb Bitboard
	b Board
	tb Board
}

func TestBoard_SetPieceBitboard(t *testing.T) {
	tests := []pieceBitboard{
		{"White Pawn", WhitePawn, Bitboard(255), Board{}, Board{wPawns: Bitboard(255)}},
		{"White Knight", WhiteKnight, Bitboard(10), Board{}, Board{wKnights: Bitboard(10)}},
		{"White Bishop", WhiteBishop, Bitboard(88445), Board{}, Board{wBishops: Bitboard(88445)}},
		{"White Rook", WhiteRook, Bitboard(10000001), Board{}, Board{wRooks: Bitboard(10000001)}},
		{"White Queen", WhiteQueen, Bitboard(1<<24), Board{}, Board{wQueens: Bitboard(1<<24)}},
		{"White King", WhiteKing, Bitboard(555556), Board{}, Board{wKing: Bitboard(555556)}},
		{"Black Pawn", BlackPawn, Bitboard(1023), Board{}, Board{bPawns: Bitboard(1023)}},
		{"Black Knight", BlackKnight, Bitboard(0), Board{}, Board{bKnights: Bitboard(0)}},
		{"Black Bishop", BlackBishop, Bitboard(65538), Board{}, Board{bBishops: Bitboard(65538)}},
		{"Black Rook", BlackRook, Bitboard(100000), Board{}, Board{bRooks: Bitboard(100000)}},
		{"Black Queen", BlackQueen, Bitboard(2), Board{}, Board{bQueens: Bitboard(2)}},
		{"Black King", BlackKing, Bitboard(1), Board{}, Board{bKing: Bitboard(1)}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T){
			board := test.b
			board.SetPieceBitBoard(test.p, test.bb)

			if !reflect.DeepEqual(board, test.tb) {
				t.Fatalf("Incorrect Bitboard: %v instead %v", test.tb, board)
			}
		})
	}
}

func TestBoard_UtilBitboards(t *testing.T) {
	testBoard := Board{
		bKing: Bitboard(1<<11),
		wKing: Bitboard(1<<45),
		bQueens: Bitboard(1<<3),
		wQueens: Bitboard(1<<60),
		bPawns: Bitboard(1<<19+1<<28),
		wPawns: Bitboard(1<<23+1<<41),
		wKnights: Bitboard(1<<24),
		bBishops: Bitboard(1<<17),
	}
	var (
		btotal = Bitboard(269_092_872)
		wtotal = Bitboard(1_152_958_888_027_357_184)
		total = Bitboard(1_152_958_888_296_450_056)
		empty = Bitboard(17_293_785_185_413_101_559)
	)
	testBoard.UtilBitboards()

	if testBoard.wPieces != wtotal {
		t.Fatalf("Incorrect calculation of white bitboard: %d instead of %d", testBoard.wPieces, wtotal)
	}
	if testBoard.bPieces != btotal {
		t.Fatalf("Incorrect calculation of black bitboard: %d instead of %d", testBoard.bPieces, btotal)
	}
	if testBoard.pieces != total {
		t.Fatalf("Incorrect calculation of total bitboard: %d instead of %d", testBoard.pieces, total)
	}
	if testBoard.empty != empty {
		t.Fatalf("Incorrect calculation of empty bitboard: %d instead of %d", testBoard.empty, empty)
	}
}