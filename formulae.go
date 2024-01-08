package main

func sum(operands []float64) float64 {
	total := 0.0
	for _, f := range operands {
		total += f
	}
	return total
}

func product(operands []float64) float64 {
	total := 1.0
	for _, f := range operands {
		total *= f
	}
	return total
}

func max(operands []float64) float64 {
	max := operands[0]
	for _, f := range operands {
		if f > max {
			max = f
		}
	}
	return max
}

func min(operands []float64) float64 {
	min := operands[0]
	for _, f := range operands {
		if f < min {
			min = f
		}
	}
	return min
}

func average(operands []float64) float64 {
	return sum(operands) / float64(len(operands))
}

func count(operands []float64) float64 {
	return float64(len(operands))
}
