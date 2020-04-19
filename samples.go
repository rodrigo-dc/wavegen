package wavegen

import "fmt"

func createSampleSlice(frequency float64, sampleRate uint) ([]float64, error) {
	wavePeriod := 1 / frequency
	samplePeriod := 1 / float64(sampleRate)
	samplesPerPeriod := int(wavePeriod / samplePeriod)

	if samplesPerPeriod < 2 {
		err := fmt.Errorf("frequency too high for the sample rate")
		return nil, err
	}

	samples := make([]float64, 0, samplesPerPeriod)

	return samples, nil
}
