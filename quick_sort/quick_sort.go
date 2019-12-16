package quick_sort

func ascending(values []int, left int, right int) {
	if left < right {
		key := values[(left+right)/2]
		i := left
		j := right

		for {
			if values[i] < key {
				i++
			}
			if values[j] > key {
				j--
			}
			if i >= j {
				break
			}
			values[i], values[j] = values[j], values[i]
		}

		ascending(values, left, i-1)
		ascending(values, j+1, right)
	}
}
