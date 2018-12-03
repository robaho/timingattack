package timingattack

import "testing"

func BenchmarkTimeNone(b *testing.B) {
	s := "This is a string"

	for i:=0;i<b.N;i++ {
		if s=="Won't Ever Match" {
			return
		}
	}
}

func BenchmarkTimeOne(b *testing.B) {
	s := "Will this match?"

	for i:=0;i<b.N;i++ {
		if s=="Won't Ever Match" {
			return
		}
	}
}