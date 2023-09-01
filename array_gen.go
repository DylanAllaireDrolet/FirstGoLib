package FirstGoLib

import "math/rand"

func ArrayGenerate(size int) (int, int, int) {

	array := make([]int, size)

	var min, max, moy int

	for i := 0; i < size; i++ {
		array[i] = rand.Intn(10000)
	}
	min = array[0]
	max = array[0]

	for _, n := range array {
		if max < n {
			max = n
		}
		if min > n {
			min = n
		}
		moy += n
	}

	return min, max, (moy / size)
}
