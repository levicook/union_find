package union_find

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Benchmark_QuickFind_Find_Identity_10000(b *testing.B) {
	bench_Find_Identity(b, NewQuickFind(10000))
}

func Benchmark_QuickFind_Find_Identity_100000(b *testing.B) {
	bench_Find_Identity(b, NewQuickFind(100000))
}

func Benchmark_QuickFind_Find_Identity_1000000(b *testing.B) {
	bench_Find_Identity(b, NewQuickFind(1000000))
}

func Benchmark_QuickFind_Union_All_1000(b *testing.B) {
	bench_Union_All(b, NewQuickFind(1000))
}

//func Benchmark_QuickFind_Union_All_10000(b *testing.B) {
//	bench_Union_All(b, NewQuickFind(10000))
//}

func Test_QuickFind(t *testing.T) {
	qf := NewQuickFind(10)

	for i := 0; i < qf.Len(); i++ {
		for j := 0; j < qf.Len(); j++ {

			if i == j {
				require.True(t, qf.Find(i, j))
			} else {
				require.False(t, qf.Find(i, j))
			}
		}
	}

	qf.Union(4, 3)
	require.True(t, qf.Find(4, 3))
	require.True(t, qf.Find(3, 4))

	qf.Union(3, 8)
	require.True(t, qf.Find(3, 8))
	require.True(t, qf.Find(8, 3))
	require.True(t, qf.Find(4, 8))
	require.True(t, qf.Find(8, 4))
}
