package chess

const (
	ranks = "87654321"
	files = "abcdefgh"
)

// we just want to store 0-63
// e.g.: Square a1 <==> 56
type Square uint8

// a1 <==> 56
// 56 % 8 = 0 ==> File(0) ==> FileA
func (s Square) File() File {
	return File(int(s) % 8)
}

// a1 <==> 56
// 56 / 8 = 7 ==> Rank(7) ==> 1
func (s Square) Rank() Rank {
	return Rank(int(s) / 8)
}

func (s Square) Color() Color {
	return Color(int(s%8+s/8)%2)
}

func NewSquare(r Rank, f File) Square {
	return Square(8*uint8(r)+uint8(f))
}

func (s Square) String() string {
	return s.File().String() + s.Rank().String()
} 

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
	invalid
)

type Rank uint8

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

func (r Rank) String() string {
	return ranks[r:r+1]
}

type File uint8

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

func (f File) String() string {
	return files[f:f+1]
}

type Color uint8

const (
	White Color = iota
	Black
	Neither
)

func (c Color) String() string {
	switch c {
	case White:
		return "white"
	case Black:
		return "black"
	}
	return ""
}