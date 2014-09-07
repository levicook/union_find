package union_find

import "testing"

func bench_Find_Identity(b *testing.B, uf UnionFind) {
	for i := 0; i < uf.Len(); i++ {
		uf.Find(i, i)
	}
}

func bench_Union_All(b *testing.B, uf UnionFind) {
	for i := 0; i < uf.Len(); i++ {
		for j := 0; j < uf.Len(); j++ {
			uf.Union(i, j)
		}
	}
}
