package chess

import (
	"fmt"
	"testing"
)

type squareTest struct {
	s Square
	r Rank
	f File
	c Color
}

var tests = []squareTest{
	{A1,Rank1,FileA,Black},
	{C7,Rank7,FileC,Black},
	{H8,Rank8,FileH,Black},
	{E3,Rank3,FileE,Black},
	{F5,Rank5,FileF,White},
}

// no tests for invalid squares yet ?
func TestSquare_Extraction(t *testing.T) {
	for i, test := range tests {
		t.Run(fmt.Sprintf("ExtractionTest %d",i), func(t *testing.T){
			rank := test.s.Rank()
			file := test.s.File()

			if rank != test.r || file != test.f {
				t.Fatalf("Wrong Square extracted: Got %v -> %s%s",test.s.String(),file.String(),rank.String())
			}
		})
	}
}

func TestNewSquare (t *testing.T) {
	for i, test := range tests {
		t.Run(fmt.Sprintf("CreationTest %d",i), func(t *testing.T){
			square := NewSquare(test.r, test.f)

			if square != test.s {
				t.Fatalf("Wrong Square generated: Got %s%s->%s != %s",test.f.String(),test.r.String(),square.String(),test.s.String())
			}
		})
	}
}

func TestSquare_Color (t *testing.T) {
	for i, test := range tests {
		t.Run(fmt.Sprintf("ColorTest %d",i), func(t *testing.T){
			square := NewSquare(test.r, test.f)

			if square.Color() != test.c {
				t.Fatalf("Wrong Color generated: Got %s->%s != %s",test.s,square.Color().String(),test.c.String())
			}
		})
	}
}

func TestStringer (t *testing.T) {
	t.Run("White Square", func(t *testing.T){
		whiteSquare := NewSquare(Rank4,FileE)
		if whiteSquare.String() != "e4" {
			t.Fatalf("Wrong Square-String generated: Got %v->%s",whiteSquare,whiteSquare.String())
		}
		if whiteSquare.Rank().String() != "4" {
			t.Fatalf("Wrong Rank-String generated: Got %v->%s",whiteSquare.Rank(),whiteSquare.Rank().String())
		}
		if whiteSquare.File().String() != "e" {
			t.Fatalf("Wrong File-String generated: Got %v->%s",whiteSquare.File(),whiteSquare.File().String())
		}
		if whiteSquare.Color().String() != "white" {
			t.Fatalf("Wrong Color-String generated: Got %v->%s",whiteSquare.Color(),whiteSquare.Color().String())
		}
	})
	t.Run("Black Square", func(t *testing.T){
		blackSquare := NewSquare(Rank7,FileG)
		if blackSquare.String() != "g7" {
			t.Fatalf("Wrong Square-String generated: Got %v->%s",blackSquare,blackSquare.String())
		}
		if blackSquare.Rank().String() != "7" {
			t.Fatalf("Wrong Rank-String generated: Got %v->%s",blackSquare.Rank(),blackSquare.Rank().String())
		}
		if blackSquare.File().String() != "g" {
			t.Fatalf("Wrong File-String generated: Got %v->%s",blackSquare.File(),blackSquare.File().String())
		}
		if blackSquare.Color().String() != "black" {
			t.Fatalf("Wrong Color-String generated: Got %v->%s",blackSquare.Color(),blackSquare.Color().String())
		}
	})
}