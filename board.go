package chess

// Board is a collection of Bitboard representations:
// For each Piece, Color etc. those are saved in corresponding numbers
type Board struct {
	wKing    Bitboard
	wQueens  Bitboard
	wRooks   Bitboard
	wBishops Bitboard
	wKnights Bitboard
	wPawns   Bitboard
	bKing    Bitboard
	bQueens  Bitboard
	bRooks   Bitboard
	bBishops Bitboard
	bKnights Bitboard
	bPawns   Bitboard
	wPieces  Bitboard
	bPieces  Bitboard
	empty    Bitboard
	pieces   Bitboard
}

func FromMap(m map[Square]Piece) *Board {
	b := &Board{}
	for _, p := range pieces {
		setSquares := make(map[Square]bool)
		for square, piece := range m {
			if piece == p {
				setSquares[square] = true
			}
		}
		b.SetPieceBitBoard(p, MakeBitboard(setSquares))
	}
	b.UtilBitboards()
	return b
}

func (b *Board) SetPieceBitBoard(piece Piece, bitboard Bitboard) {
	switch piece {
	case WhitePawn:
		b.wPawns = bitboard
	case WhiteKnight:
		b.wKnights = bitboard
	case WhiteBishop:
		b.wBishops = bitboard
	case WhiteRook:
		b.wRooks = bitboard
	case WhiteQueen:
		b.wQueens = bitboard
	case WhiteKing:
		b.wKing = bitboard
	case BlackPawn:
		b.bPawns = bitboard
	case BlackKnight:
		b.bKnights = bitboard
	case BlackBishop:
		b.bBishops = bitboard
	case BlackRook:
		b.bRooks = bitboard
	case BlackQueen:
		b.bQueens = bitboard
	case BlackKing:
		b.bKing = bitboard
	}
}

func (b *Board) UtilBitboards() {
	b.wPieces = b.wPawns + b.wKnights + b.wBishops + b.wRooks + b.wQueens + b.wKing
	b.bPieces = b.bPawns + b.bKnights + b.bBishops + b.bRooks + b.bQueens + b.bKing
	b.pieces = b.wPieces + b.bPieces
	b.empty = ^b.pieces
}
