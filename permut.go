package main

import (
	"flag"
	"fmt"
	"github.com/ivanmilov/fact"
	"sort"
)

type arr []int

func (g arr) Len() int           { return len(g) }
func (h arr) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h arr) Less(i, j int) bool { return h[i] < h[j] }

func main() {

	numbPtr := flag.Int("n", 4, "an int")
	flag.Parse()
	var N = int(*numbPtr)

	s := make([]int, N)
	// gen first {1,2,3,4,...,N}
	for i := 1; i <= N; i++ {
		s[i-1] = i
	}

	nf := fact.FactTree(int64(N)).Int64()
	for i := int64(0); i < nf; i++ {
		print(s)
		s = get(s)
	}
}

func print(s arr) {
	for _, i := range s {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

func min(s arr) int {
	indx := 0
	for i := 1; i < len(s); i++ {
		if s[i] < s[indx] {
			indx = i
		}
	}
	return indx
}

func getNextRightIndex(el, def, addinex int, s arr) int {
	for i, c := range s {
		if c == el+1 {
			return i + addinex
		}
	}
	return min(s) + addinex
}

func getRightSlice(part arr) arr {
	// find next to el; part[0] = el
	el := part[0]
	p := part[1:]

	if len(p) == 1 {
		el, p[0] = p[0], el
	} else {
		// find next bigger than el
		sort.Sort(p)
		for i := 0; i < len(p); i++ {
			if p[i] > el {
				el, p[i] = p[i], el
				break
			}
		}
	}

	a := []int{el}
	a = append(a, p...)
	return a
}

func get(s arr) arr {
	for i := len(s) - 1; i > 0; i-- {
		r := s[i]
		l := s[i-1]
		if r > l {
			s = append(s[:i-1], getRightSlice(s[i-1:])...)
			return s
		}
	}
	return s
}
