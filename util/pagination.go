package util

import "math"

func ConvertOffsetToPage(offset, limit int) int {
	if limit < 1 {
		limit = 10
	}
	page := (offset / limit) + 1
	return page
}

func CalculateMaxPage(totalRowCount, limit int) int {
	if limit <= 0 {
		return 0 // Hindari pembagian oleh nol
	}
	return int(math.Ceil(float64(totalRowCount) / float64(limit)))
}
