package union_find

func NewQuickUnion(size int) UnionFind {
	return newQuickUnion(newIds(size))
}

func newQuickUnion(ids []int) *quickUnion {
	qu := new(quickUnion)
	qu.ids = ids
	return qu
}

type quickUnion struct{ ids []int }

func (qu *quickUnion) Len() int { return len(qu.ids) }

func (qu *quickUnion) root(i int) int {
	for i != qu.ids[i] {
		i = qu.ids[i]
	}
	return i
}

func (qu *quickUnion) Find(p, q int) bool {
	return qu.root(p) == qu.root(q)
}

func (qu *quickUnion) Union(p, q int) {
	i, j := qu.root(p), qu.root(q)
	qu.ids[i] = j
}
