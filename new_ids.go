package union_find

func newIds(size int) []int {
	ids := make([]int, size)
	for i := range ids {
		ids[i] = i
	}
	return ids
}
