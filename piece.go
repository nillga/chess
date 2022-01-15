package chess

import "strings"

// AbstractPiece is an integer representation for the Chess-Pieces' Types
type AbstractPiece uint8

var (
	abstractChars = []string{"", "P", "N", "B", "R", "Q", "K"}
	pieceUnicodes = []string{" ", "♙", "♘", "♗", "♖", "♕", "♔", "♟", "♞", "♝", "♜", "♛", "♚"}
)

// AbstractPiece is implemented with an enumeration
const (
	None AbstractPiece = iota
	Pawn
	Knight
	Bishop
	Rook
	Queen
	King
)

// String -> Stringer interface
func (a AbstractPiece) String() string {
	return abstractChars[a]
}

// The Piece type allows an integer representation of a chess Piece
type Piece uint8

// Color returns the explicit Piece's color
func (p Piece) Color() Color {
	if p == Empty {
		return Neither
	}
	if p > WhiteKing {
		return Black
	}
	return White
}

// Abstract returns the AbstractPiece from which the given explicit Piece has been derived
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

// String fulfills the Stringer interface
func (p Piece) String() string {
	return pieceUnicodes[p]
}

// MakePiece combines the given AbstractPiece and Color to return an explicit Piece
func MakePiece(a AbstractPiece, c Color) Piece {
	p := Piece(a)

	if p == Empty {
		return p
	}
	return p + 6*Piece(c)
}

// An enumeration containing all explicit chess pieces
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

// FEN returns the FEN-notation-styled symbol for the given Piece
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
