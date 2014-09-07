package union_find

type UnionFind interface {
	Find(int, int) bool
	Union(int, int)
	Len() int
}
