package version

import (
	"strconv"
	"strings"
)

// TODO: sort versions (interface)

type suffix struct {
	txt string
	rev uint
}

type PMSVersion struct {
	nums []uint // array of number versions 10.0.103 -> [10 0 103]
	// TODO implement: letter string    // char e.g. Openssl 1.0.1j -> j
	// TODO implement: suf    []suffix  // suffix e.g. _alpha123 -> { "alpha", "123" }
	// TODO implement: rel    string    // release e.g. -r1 -> "1" (if none 0 is implied)
}

func parse(v string) (p PMSVersion, ok bool) {
	// TODO: `for until number` and then split pass only numbers with . to parsenums
	if v == "" {
		return
	}
	/*/ parseInts{
	i := 0
	for i < len(v) && isNum(v[i:i]) || v[i] == '.' {
		i++
	}
	sNums := v[:i]
	*/                   // } parseInts
	nums := parseNums(v) //sNums
	return PMSVersion{nums}, true
}

// compare returns int depending on which PMSVersion is higher
//   0 if A == B
//   1 if A >  B
//  -1 if A <  B
func compare(v1, v2 PMSVersion) int {
	return compareNums(v1.nums, v2.nums)
}

func parseNums(v string) []uint {
	sNums := strings.Split(v, ".")
	var nums []uint
	for _, s := range sNums {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		nums = append(nums, uint(i))
	}
	return nums
}

// compareNums returns int depending on which version is higher
//   0 if A == B
//   1 if A >  B
//  -1 if A <  B
func compareNums(A, B []uint) int {
	Ann, Bnn := len(A), len(B)
	var shorter []uint
	if Ann < Bnn {
		shorter = A
	} else {
		shorter = B
	}
	// Algorithm 3.2
	if A[0] > B[0] {
		return 1
	} else if A[0] < B[0] {
		return -1
	} else {
		// Algorithm 3.3
		for i, _ := range shorter {
			// omitted stringwise comparison because it will be redundant? (I think)
			if A[i] > B[i] {
				return 1
			} else if A[i] < B[i] {
				return -1
			}
		}
	}
	// I think this should be earlier because it will resolve quicker
	if Ann > Bnn {
		return 1
	} else if Ann < Bnn {
		return -1
	}
	return 0
}

func isNum(v string) bool {
	i := 0
	for i < len(v) && '0' <= v[i] && v[i] <= '9' {
		i++
	}
	return i == len(v)
}
