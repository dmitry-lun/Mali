package parser

import "github.com/dmitry-lun/Mali/pkg/entropy"

func CalculateEntropy(data []byte) float64 {
	return entropy.Calculate(data)
}
