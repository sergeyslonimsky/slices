package slices

// ArrWalk creates a new array populated with the results of calling a provided function
// on every element in the calling array.
func ArrWalk[I, T any](arr []I, callback func(key int, value I) T) []T {
	r := make([]T, 0, len(arr))

	for i, v := range arr {
		r = append(r, callback(i, v))
	}

	return r
}

// ArrForEach executes a provided function once for each array element.
func ArrForEach[I any](arr []I, callback func(key int, value I)) {
	for i, v := range arr {
		callback(i, v)
	}
}

// ArrFilter creates a copy of a given array,
// filtered down to just the elements from the given array that pass the test implemented by the provided function.
func ArrFilter[I any](arr []I, callback func(key int, value I) bool) []I {
	r := make([]I, 0, len(arr))

	for i, v := range arr {
		if callback(i, v) {
			r = append(r, v)
		}
	}

	return r
}

// ArrConcat creates a new array with all the elements of the provided arrays.
func ArrConcat[I any](arr ...[]I) []I {
	r := make([]I, 0, len(arr[0]))

	for _, arrN := range arr {
		for _, v := range arrN {
			r = append(r, v)
		}
	}

	return r
}

// ArrEvery tests whether all elements in the array pass the test implemented by the provided function.
func ArrEvery[I any](arr []I, callback func(value I) bool) bool {
	if len(arr) == 0 {
		return true
	}

	for _, v := range arr {
		if !callback(v) {
			return false
		}
	}

	return true
}

// ArrUniq creates a new array with all unique values from provided array.
// Provided array should contain comparable values.
// For non-comparable values use ArrHashUniq.
func ArrUniq[I comparable](arr []I) []I {
	r := make([]I, 0, len(arr))
	uniqMap := make(map[I]bool, len(arr))

	for _, v := range arr {
		uniqMap[v] = true
	}

	for v := range uniqMap {
		r = append(r, v)
	}

	return r
}

// ArrHashUniq creates a new array with all unique values comparable by provided hashFunc.
// HashFunc should return unique string for every unique element.
func ArrHashUniq[I any](arr []I, hashFunc func(value I) string) []I {
	r := make([]I, 0, len(arr))
	uniqMap := make(map[string]I, len(arr))

	for _, v := range arr {
		uniqMap[hashFunc(v)] = v
	}

	for _, v := range uniqMap {
		r = append(r, v)
	}

	return r
}

// ArrFind returns the first element in the provided array that satisfies the provided testing function.
// If no values satisfy the testing function, empty value and false are returned.
func ArrFind[I any](arr []I, callback func(key int, value I) bool) (I, bool) {
	for i, v := range arr {
		if callback(i, v) {
			return v, true
		}
	}

	return *new(I), false
}

// ArrFindIndex returns the index of the first element in the provided array
// that satisfies the provided testing function.
// If no values satisfy the testing function, -1 and false are returned.
func ArrFindIndex[I any](arr []I, callback func(key int, value I) bool) (int, bool) {
	for i, v := range arr {
		if callback(i, v) {
			return i, true
		}
	}

	return -1, false
}

// ArrReverse reverses an array in place and returns the new array.
// The first array element now becoming the last, and the last array element becoming the first.
func ArrReverse[I any](arr []I) []I {
	res := make([]I, len(arr))
	copy(res, arr)

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return res
}
