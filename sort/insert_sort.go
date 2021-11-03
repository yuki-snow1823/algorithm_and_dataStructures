package InsertionSort

// InsertionSort 挿入ソート
func InsertionSort(a []int) {
	var (
		i, j, key int
	)

	for j = 1; j < len(a); j++ {
		key = a[j]
		i = j - 1
		
		for i >= 0 && a[i] > key {
			a[i+1] = a[i]
			i--
		}
		a[i+1] = key
	}
	return
}
