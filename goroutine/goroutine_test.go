package goroutine

import (
	"testing"
	"time"
)

func TestPrintNon(t *testing.T) {
	PrintNon()
}

func TestPrintGo(t *testing.T) {
	PrintGo()
	time.Sleep(1 * time.Second)
}

func BenchmarkPrintNon(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrintNon()
	}
}

func BenchmarkPrintGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrintGo()
	}
}
