package bit

func HammingDistance(x int, y int) int {
	//return bits.OnesCount(uint(x ^ y))
	count := 0
	for r := uint(x ^ y); r > 0; r >>= 1 {
		if r&1 == 1 {
			count++
		}
	}
	return count
}
