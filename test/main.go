package main

import (
	"fmt"

	permutations "github.com/therox/go-permutations"
)

type testStruct struct {
	ID   int
	Name string
}

type testStructSlice []testStruct

func main() {
	a := testStructSlice{
		testStruct{
			3, "Player 3",
		},
		testStruct{
			5, "Player 5",
		},
		testStruct{
			2, "Player 2",
		},
		testStruct{
			1, "Player 1",
		},
		testStruct{
			4, "Player 4",
		},
	}
	counter := 1
	permutations.NewPermutation(a)
	for {
		fmt.Println(counter, ": ", a[:4])
		counter++
		if !permutations.NextKPermutation(a, 4) {
			break
		}
	}

}

func (is testStructSlice) Len() int {
	return len(is)

}

// Less reports whether the element with
// index i should sort before the element with index j.
func (is testStructSlice) Less(i, j int) bool {
	return is[i].ID < is[j].ID

}

// Swap swaps the elements with indexes i and j.
func (is testStructSlice) Swap(i, j int) {
	is[j], is[i] = is[i], is[j]
}
