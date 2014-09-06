package union_find

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_QuickFind(t *testing.T) {
	size := 10

	uf := NewQuickFind(size)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {

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
