package union_find

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Benchmark_QuickUnion_Find_Identity_10000(b *testing.B) {
	bench_Find_Identity(b, NewQuickUnion(10000))
}

func Benchmark_QuickUnion_Find_Identity_100000(b *testing.B) {
	bench_Find_Identity(b, NewQuickUnion(100000))
}

func Benchmark_QuickUnion_Find_Identity_1000000(b *testing.B) {
	bench_Find_Identity(b, NewQuickUnion(1000000))
}

func Benchmark_QuickUnion_Union_All_1000(b *testing.B) {
	bench_Union_All(b, NewQuickUnion(1000))
}

func Benchmark_QuickUnion_Union_All_10000(b *testing.B) {
	bench_Union_All(b, NewQuickUnion(10000))
}

func Test_QuickUnion(t *testing.T) {
	uf := NewQuickUnion(10)

	for i := 0; i < uf.Len(); i++ {
		for j := 0; j < uf.Len(); j++ {

			if i == j {
				require.True(t, uf.Find(i, j))
			} else {
				require.False(t, uf.Find(i, j))
			}
		}
	}

	uf.Union(4, 3)
	require.True(t, uf.Find(4, 3))
	require.True(t, uf.Find(3, 4))

	uf.Union(3, 8)
	require.True(t, uf.Find(3, 8))
	require.True(t, uf.Find(8, 3))
	require.True(t, uf.Find(4, 8))
	require.True(t, uf.Find(8, 4))
}
