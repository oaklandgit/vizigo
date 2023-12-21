package main

func Sum(operands []float64) float64 {
    total := 0.0
    for _, f := range operands {
        total += f
    }
    return total
}

func Product(operands []float64) float64 {
	total := 1.0
	for _, f := range operands {
		total *= f
	}
	return total
}

func Max(operands []float64) float64 {
	max := operands[0]
	for _, f := range operands {
		if f > max {
			max = f
		}
	}
	return max
}

func Min(operands []float64) float64 {
	min := operands[0]
	for _, f := range operands {
		if f < min {
			min = f
		}
	}
	return min
}

func Average(operands []float64) float64 {
	return Sum(operands) / float64(len(operands))
}

func Count(operands []float64) float64 {
	return float64(len(operands))
}