package permutations

import (
	"sort"
)

// NewKPermutation generates initial permutation
func NewPermutation(data sort.Interface) {
	sort.Sort(data)
}

// NextPermutation generate next permutation
func NextPermutation(data sort.Interface) bool {
	return NextKPermutation(data, data.Len())
}

// NextKPermutation generate next permutation
func NextKPermutation(data sort.Interface, k int) bool {
	n := data.Len()

	var edge, j, i int
	edge = k - 1

	// find j in (k…n-1) where a[j] > a[edge]
	j = k

	for (j < n) && (data.Less(j, edge)) {
		j++
	}

	if j < n {
		// swap a[egde], a[j]
		data.Swap(edge, j)
	} else {
		// reverse a[k] to a[n-1]
		for i, j = k, n-1; i < j; i, j = i+1, j-1 {
			data.Swap(i, j)

		}

		// find rightmost ascent to left of edge
		i = edge - 1
		for i >= 0 && data.Less(i+1, i) {
			i--
		}

		if i < 0 {
			return false // no more permutations
		}

		// find j in (n-1…i+1) where aj > ai
		j = n - 1
		for j > i && data.Less(j, i) {
			j--
		}
		// swap a[i], a[j]
		data.Swap(i, j)

		// reverse a[i+1] to an-1
		for i, j = i+1, n-1; i < j; i, j = i+1, j-1 {
			data.Swap(i, j)
		}
	}
	return true
}
