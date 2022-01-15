package chess

import (
	"fmt"
	"strconv"
	"strings"
)

// TODO: parse FEN into position & vice-versa -> position struct!

/*
	FEN (Forsyth-Edwards-Notation is used to represent a single position as a string

	Starting Position in FEN:					Turns since last take/pawn move
														 v
	rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1  <- Next ply
												^	^  L en passant
												|	L castling options
												|
											Turn Player
	♜♞♝♛♚♝♞♜
	♟♟♟♟♟♟♟♟




	♙♙♙♙♙♙♙♙
	♖♘♗♕♔♗♘♖
*/

// The FEN type allows FEN-notation related actions on a string, representing FEN code
type FEN string

/*
Parse() function to parse entire FEN notation into a **Position**
func (f FEN) Parse() (*Board, error) {
	chunks := strings.Split(string(f), " ")
	// chunks is {"Board Rep","Turn Player", "Castling", "ep", "50ply", "nextply"}
	return BoardFromFEN(chunks[0])
}
*/

// BoardFromFEN takes the board-representation part of a FEN string and turns it into a *Board
// it returns an error, when the Board structure is invalid: invalid symbols, wrong rank/file-counts
func BoardFromFEN(fen string) (*Board, error) {
	m := make(map[Square]Piece)
	ranks := strings.Split(fen, "/")
	if l := len(ranks); l != 8 {
		return nil, fmt.Errorf("wrong rank number: %d", l)
	}
	for i, rankFEN := range ranks {
		rank := Rank(i)
		rankMap, err := decodeRank(rank, rankFEN)
		if err != nil {
			return nil, err
		}
		for file, piece := range rankMap {
			m[NewSquare(rank, file)] = piece
		}
	}

	return FromMap(m), nil
}

func decodeRank(r Rank, s string) (map[File]Piece, error) {
	m := make(map[File]Piece)
	file := 0
	for _, val := range s {
		valString := fmt.Sprintf("%c", val)
		thisPiece := fenPieceMap[valString]
		if thisPiece == Empty {
			blank, err := strconv.Atoi(valString)
			if err != nil {
				return nil, err
			}
			file += blank
			continue
		}
		m[File(file)] = thisPiece
		file++
	}
	if file != 8 {
		return nil, fmt.Errorf("incorrect File Count %d in Rank %s", file, r.String())
	}
	return m, nil
}

var fenPieceMap = map[string]Piece{
	"k": BlackKing, "K": WhiteKing,
	"q": BlackQueen, "Q": WhiteQueen,
	"r": BlackRook, "R": WhiteRook,
	"b": BlackBishop, "B": WhiteBishop,
	"n": BlackKnight, "N": WhiteKnight,
	"p": BlackPawn, "P": WhitePawn,
}
