package chess

import (
	"fmt"
	"testing"
)

type bitboardTest struct {
	b Bitboard
	s Square
	r bool
}

var bitboardTests = []bitboardTest{
	{Bitboard(1<<40), A3, true},
	{Bitboard(1<<25), A3, false},
	{Bitboard(1), A8, true},
	{Bitboard(1<<62),G1, true},
	{Bitboard(68000),A6, true},
}

func TestBitboard_IsSet(t *testing.T) {
	for i, test := range bitboardTests {
		t.Run(fmt.Sprintf("Bitboard_Occupation %d",i), func(t *testing.T) {
			if set :=test.b.IsSet(test.s); set != test.r {
				t.Fatalf("Bitboard %b | %d has wrongly checked Square %s: got: %v, want: %v", test.b, test.b, test.s.String(), set, test.r)
			}
		})
	}
}

func TestBitboard_areSet(t *testing.T) {
	b := Bitboard(68000)
	t.Run("They're set", func(t *testing.T) {
		if !b.areSet(A6,D7,A7,H8,F8) {
			t.Fatalf("Bitboard %b | %d does not contain the set squares!", b,b)
		}
	})
	t.Run("They're not", func(t *testing.T) {
		if b.areSet(A6,D7,A7,H8,F8,F4) {
			t.Fatalf("Bitboard %b | %d does contain the non-set squares!", b,b)
		}
	})
}

func TestMakeBitboard(t *testing.T) {
	testMap := map[Square]bool{
		A8: true, H7: true, B5: true, G4: true, C2: true, F1: true, E3: true, D6: true,
	}
	bitBoard := MakeBitboard(testMap)
	if !bitBoard.areSet(A8,H7,B5,G4,C2,F1,E3,D6) {
		t.Fatalf("Bitboard was not calculated properly!")
	}
}
