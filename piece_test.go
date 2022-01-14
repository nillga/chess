package chess

import (
	"fmt"
	"testing"
)

type pieceTest struct {
	a AbstractPiece
	c Color
	p Piece
}

func TestPiece(t *testing.T) {
	tests := []pieceTest{
		{Pawn, White, WhitePawn},
		{Knight, Black, BlackKnight},
		{Queen, Black, BlackQueen},
		{King, White, WhiteKing},
		{None, Neither, Empty},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Piece_Color %d", i), func(t *testing.T) {
			if test.p.Color() != test.c {
				t.Fatalf("Color is not detected correctly! %s -> %s", test.c.String(), test.p.Color().String())
			}
		})
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("Piece_Abstract %d", i), func(t *testing.T) {
			if test.p.Abstract() != test.a {
				t.Fatalf("Type is not detected correctly! %s -> %s", test.a.String(), test.p.Abstract().String())
			}
		})
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("MakePiece %d", i), func(t *testing.T) {
			piece := MakePiece(test.a, test.c)
			if piece != test.p {
				t.Fatalf("Piece is not made correctly! %s -> %s", test.p.String(), piece.String())
			}
		})
	}
}

func TestStringers(t *testing.T) {
	blackPiece, whitePiece, noPiece := BlackRook, WhiteBishop, Empty

	if blackPiece.String() != "♜" {
		t.Fatalf("Incorrect string %s instead of ♜", blackPiece.String())
	}
	if whitePiece.String() != "♗" {
		t.Fatalf("Incorrect string %s instead of ♗", whitePiece.String())
	}
	if noPiece.String() != " " {
		t.Fatalf("Incorrect string %s instead of blank", noPiece.String())
	}
	if blackPiece.Color().String() != "black" {
		t.Fatalf("Incorrect string %s instead of 'black'", blackPiece.Color().String())
	}
	if whitePiece.Color().String() != "white" {
		t.Fatalf("Incorrect string %s instead of 'white'", whitePiece.Color().String())
	}
	if noPiece.Color().String() != "" {
		t.Fatalf("Incorrect string %s instead of ''", noPiece.Color().String())
	}
	if blackPiece.Abstract().String() != "R" {
		t.Fatalf("Incorrect string %s instead of 'R'", blackPiece.Abstract().String())
	}
	if whitePiece.Abstract().String() != "B" {
		t.Fatalf("Incorrect string %s instead of 'B'", whitePiece.Abstract().String())
	}
	if noPiece.Abstract().String() != "" {
		t.Fatalf("Incorrect string %s instead of ''", noPiece.Abstract().String())
	}
	if blackPiece.FEN() != "r" {
		t.Fatalf("Incorrect FEN char %s instead of 'r'", blackPiece.FEN())
	}
	if whitePiece.FEN() != "B" {
		t.Fatalf("Incorrect FEN char %s instead of 'B'", whitePiece.FEN())
	}
	if noPiece.FEN() != "" {
		t.Fatalf("Incorrect FEN char %s instead of ''", noPiece.FEN())
	}
}