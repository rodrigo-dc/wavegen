package wavegen

import (
	"math"
)

// Sine returns the samples of a sinusoid with values ranging from
// -1.0 to 1.0. The phase is 0.
func Sine(frequency float64, sampleRate uint) ([]float64, error) {
	samples, err := createSampleSlice(frequency, sampleRate)
	if err != nil {
		return nil, err
	}

	numOfSamples := cap(samples)

	sampleStep := (math.Pi * 2) / float64(numOfSamples)

	phi := 0.0
	for s := 0; s < numOfSamples; s++ {
		samples = append(samples, math.Sin(phi))
		phi += sampleStep
	}

	return samples, nil
}
