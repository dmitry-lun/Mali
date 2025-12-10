package entropy

import "math"

func Calculate(data []byte) float64 {
	if len(data) == 0 {
		return 0
	}
	var freq [256]int
	for _, b := range data {
		freq[b]++
	}
	entropy := 0.0
	for _, f := range freq {
		if f == 0 {
			continue
		}
		p := float64(f) / float64(len(data))
		entropy += -p * math.Log2(p)
	}
	return entropy
}
