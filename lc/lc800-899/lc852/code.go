func peakIndexInMountainArray(arr []int) int {
	i := 0
	j := len(arr) - 1
	var p int

	for p = (i + j) / 2; arr[p] < arr[p+1] || arr[p] < arr[p-1]; p = (i + j) / 2 {

		if arr[p] > arr[p+1] {
			j = p
		} else {
			i = p
		}

	}

	return p
}