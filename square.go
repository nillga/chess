package chess

// TODO: limit exportation

const (
	ranks = "87654321"
	files = "abcdefgh"
)

// Square is the numeric representation of a square
// the squares are enumerated from 0-63, starting at A8 via H8 to H1
// This coding allows simple mathematic operation to access a Square's core attributes or description,
// namely their coordinates and their color:
// The Rank is extracted by 8 - Square / 8, as we count against the enumeration
// The File is extracted by Square % 8, pretty simple
// The Color is evaluated by checking for the characteristics of both Rank and File:
// If the sum of File and reverse Rank is even, the Square is White, if odd, then Black.
type Square uint8

// File extracts the File of a given Square
func (s Square) File() File {
	return File(int(s) % 8)
}

// Rank extracts the Rank of a given Square
func (s Square) Rank() Rank {
	return Rank(int(s) / 8)
}

// Color extracts the Color of the given Square
func (s Square) Color() Color {
	return Color(int(s%8+s/8) % 2)
}

// NewSquare takes a Rank and a File and returns the described Square
func NewSquare(r Rank, f File) Square {
	return Square(8*uint8(r) + uint8(f))
}

// String makes Square implement the Stringer interface
func (s Square) String() string {
	return s.File().String() + s.Rank().String()
}

// Constant mapping of Square-notation and Bitboard-representation
const (
	A8 Square = iota
	B8
	C8
	D8
	E8
	F8
	G8
	H8
	A7
	B7
	C7
	D7
	E7
	F7
	G7
	H7
	A6
	B6
	C6
	D6
	E6
	F6
	G6
	H6
	A5
	B5
	C5
	D5
	E5
	F5
	G5
	H5
	A4
	B4
	C4
	D4
	E4
	F4
	G4
	H4
	A3
	B3
	C3
	D3
	E3
	F3
	G3
	H3
	A2
	B2
	C2
	D2
	E2
	F2
	G2
	H2
	A1
	B1
	C1
	D1
	E1
	F1
	G1
	H1
)

// The Rank type allows us to convert the reverse Rank into the actual Rank
type Rank uint8

// The mapping of reverse and actual Rank
const (
	Rank8 Rank = iota
	Rank7
	Rank6
	Rank5
	Rank4
	Rank3
	Rank2
	Rank1
)

// String makes Rank fit the Stringer interface
func (r Rank) String() string {
	return ranks[r : r+1]
}

// File is an enumeration for the Files of a board
type File uint8

// That way the 0-file turns into FileA
const (
	FileA File = iota
	FileB
	FileC
	FileD
	FileE
	FileF
	FileG
	FileH
)

// String allows File to use the Stringer interface
func (f File) String() string {
	return files[f : f+1]
}

// Color type is representing a Piece' or Square's color
type Color uint8

// The 3 options, Neither is used e.g. for errors
const (
	White Color = iota
	Black
	Neither
)

// String makes Color use Stringer as well
func (c Color) String() string {
	switch c {
	case White:
		return "white"
	case Black:
		return "black"
	}
	return ""
}
