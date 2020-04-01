package constantime_test

import (
	"fmt"
	"testing"
)

func BenchmarkConstantime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("Constant Time High Performance Go")
	}
}
