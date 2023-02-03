func minibus(n uint32, m uint32) {
	pivot := math.Ceil(float64(m) / 2)
	current := pivot

	for n != 0 {
		fmt.Println(current)
		n--
		diff := pivot - current

		switch {
		case uint32(current) == m:
			current = pivot
		case diff > 0:
			current = pivot + diff
		case diff < 0:
			current = pivot + diff - 1
		default:
			current = pivot - 1
		}
	}
}
