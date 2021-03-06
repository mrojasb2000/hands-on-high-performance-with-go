package constantime

import (
	"fmt"
	"testing"
)

func BenchmarkConstantime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("Constant Time High Performance Go")
	}
}

func BenchmarkThree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ThreeWords()
	}
}

func BenchmarkTen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TenWords()
	}
}
