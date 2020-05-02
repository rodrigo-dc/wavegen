package wavegen

import "math"

func shift(samples []float64, offset float64) {
	for i := range samples {
		samples[i] += offset
	}
}

func scale(samples []float64, factor float64) {
	for i := range samples {
		samples[i] *= factor
	}
}

func to8Bits(input []float64) []uint8 {

	inputCopy := make([]float64, len(input))
	copy(inputCopy, input)
	shift(inputCopy, 1)
	scale(inputCopy, 127.5)

	output := make([]byte, len(inputCopy))

	for i := range inputCopy {
		output[i] = uint8(math.Round(inputCopy[i]))
	}

	return output
}
