package chess

import "strings"

type AbstractPiece uint8

var (
	abstractChars = []string{"", "P", "N", "B", "R", "Q", "K"}
	pieceUnicodes = []string{" ", "♙", "♘", "♗", "♖", "♕", "♔", "♟", "♞", "♝", "♜", "♛", "♚"}
)

const (
	None AbstractPiece = iota
	Pawn
	Knight
	Bishop
	Rook
	Queen
	King
)

func (a AbstractPiece) String() string {
	return abstractChars[a]
}

type Piece uint8

func (p Piece) Color() Color {
	if p == Empty {
		return Neither
	}
	if p > WhiteKing {
		return Black
	}
	return White
}

func (p Piece) Abstract() AbstractPiece {
	switch p {
	case BlackPawn, WhitePawn:
		return Pawn
	case BlackKnight, WhiteKnight:
		return Knight
	case BlackBishop, WhiteBishop:
		return Bishop
	case BlackRook, WhiteRook:
		return Rook
	case BlackQueen, WhiteQueen:
		return Queen
	case BlackKing, WhiteKing:
		return King
	default:
		return None
	}
}

func (p Piece) String() string {
	return pieceUnicodes[p]
}

func MakePiece(a AbstractPiece, c Color) Piece {
	p := Piece(a)

	if p == Empty {
		return p
	}
	return p + 6*Piece(c)
}

const (
	Empty Piece = iota
	WhitePawn
	WhiteKnight
	WhiteBishop
	WhiteRook
	WhiteQueen
	WhiteKing
	BlackPawn
	BlackKnight
	BlackBishop
	BlackRook
	BlackQueen
	BlackKing
)

func (p Piece) FEN() string {
	symbol := p.Abstract().String()
	if p.Color() == Black {
		return strings.ToLower(symbol)
	}
	return symbol
}

var pieces = []Piece{
	WhitePawn,
	WhiteKnight,
	WhiteBishop,
	WhiteRook,
	WhiteQueen,
	WhiteKing,
	BlackPawn,
	BlackKnight,
	BlackBishop,
	BlackRook,
	BlackQueen,
	BlackKing,
}
