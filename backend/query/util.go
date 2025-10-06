package query

func Intersection[T comparable](slices ...[]T) []T {
	if len(slices) == 0 {
		return nil
	}

	counts := make(map[T]int)
	result := []T{}

	// Count occurrences per element across slices
	for i, s := range slices {
		seen := make(map[T]bool)
		for _, v := range s {
			// avoid counting duplicates in same slice
			if !seen[v] {
				seen[v] = true
				counts[v]++
			}
		}
		// Optional optimization: if it's the first slice,
		// initialize all counts to 1 for its elements.
		if i == 0 {
			for k := range seen {
				counts[k] = 1
			}
		}
	}

	// Keep only those appearing in all slices
	for v, c := range counts {
		if c == len(slices) {
			result = append(result, v)
		}
	}

	return result
}

func Map[T any, U any](src []T, fn func(T) U) []U {
	result := make([]U, len(src))
	
	for i, v := range src {
		result[i] = fn(v)
	}
	
	return result
}
