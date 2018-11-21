//https://gobyexample.com/collection-functions
///https://github.com/mmcgrana/gobyexample#license

package iterator

//IndexOfInteger Returns the first index of the target int `t`, or -1 if no match is found.
func IndexOfInteger(vs []int, t int) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

//IsIntegerIncluded Returns true if the target int t is in the slice.
func IsIntegerIncluded(vs []int, t int) bool {
	return IndexOfInteger(vs, t) >= 0
}

//MatchAnyInteger Returns true if one of the integers in the slice satisfies the predicate f.
func MatchAnyInteger(vs []int, f func(int) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

//MatchAllIntegers Returns true if all of the integers in the slice satisfy the predicate f.
func MatchAllIntegers(vs []int, f func(int) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

//FilterIntegers Returns a new slice containing all integers in the slice that satisfy the predicate f.
func FilterIntegers(vs []int, f func(int) bool) []int {
	var vsf []int
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

//MapIntegers Returns a new slice containing the results of applying the function f to each int in the original slice.
func MapIntegers(vs []int, f func(int) int) []int {
	vsm := make([]int, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}
