package main

func selectionSort(ar []int) {
	n := len(ar)
	for i := 0; i <= n-2; i++ {
		min := i

		for j := i + 1; j <= n-1; j++ {
			if ar[j] < ar[min] {
				min = j
			}
		}
		swap(ar[i], ar[min])
	}
}

func swap(m, n int) {

}
