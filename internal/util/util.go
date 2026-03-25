package util

func calcOffset(page, size int) int {
	return (page - 1) * size
}
