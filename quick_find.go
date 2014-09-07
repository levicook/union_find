package union_find

func NewQuickFind(size int) UnionFind {
	return newQuickFind(newIds(size))
}

func newQuickFind(ids []int) *quickFind {
	qf := new(quickFind)
	qf.ids = ids
	return qf
}

type quickFind struct{ ids []int }

func (qf *quickFind) Len() int { return len(qf.ids) }

func (qf *quickFind) Find(p, q int) bool {
	return qf.ids[p] == qf.ids[q]
}

func (qf *quickFind) Union(p, q int) {
	pid := qf.ids[p]
	qid := qf.ids[q]

	for i := 0; i < len(qf.ids); i++ {
		if qf.ids[i] == pid {
			qf.ids[i] = qid
		}
	}
}
