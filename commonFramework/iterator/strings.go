//https://gobyexample.com/collection-functions
///https://github.com/mmcgrana/gobyexample#license

package iterator

//IndexOfString Returns the first index of the target string `t`, or -1 if no match is found.
func IndexOfString(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

//IsStringIncluded Returns true if the target string t is in the slice.
func IsStringIncluded(vs []string, t string) bool {
	return IndexOfString(vs, t) >= 0
}

//MatchAnyString Returns true if one of the strings in the slice satisfies the predicate f.
func MatchAnyString(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

//MatchAllStrings Returns true if all of the strings in the slice satisfy the predicate f.
func MatchAllStrings(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

//FilterStrings Returns a new slice containing all strings in the slice that satisfy the predicate f.
func FilterStrings(vs []string, f func(string) bool) []string {
	var vsf []string
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

//MapStrings Returns a new slice containing the results of applying the function f to each string in the original slice.
func MapStrings(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}
