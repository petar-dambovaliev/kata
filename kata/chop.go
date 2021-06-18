package kata

// chop is a iterative binary search implementation
// it is the only implementation that makes sense for this language
// approaches using lambdas and/or recursion will be inefficient
// because Go doesn't have tail recursion optimization in most cases
// and each lambda needs its own stack
// in general, recursion should be used only in algorithms where such would
// make the code simpler or an iterative solution is not possible
// it finds needle in the haystack and returns its index or where it should be, if it doesn't
func Chop(needle int, haystack []int) int {
	i := 0
	length := len(haystack)

	for i < length {
		//get the middle avoid overflow
		h := int(uint(i+length) >> 1)
		// if the current middle element is less
		// go to the right
		if haystack[h] < needle {
			i = h + 1
			// otherwise go to the left
		} else {
			length = h
		}
	}

	return i
}