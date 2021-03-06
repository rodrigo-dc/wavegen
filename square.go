package wavegen

// Square returns the samples of a square wave with values ranging from
// -1.0 to 1.0. The duty cycle is 50%.
func Square(frequency float64, sampleRate uint) ([]float64, error) {
	samples, err := createSampleSlice(frequency, sampleRate)
	if err != nil {
		return nil, err
	}

	numOfSamples := cap(samples)

	for s := 0; s < numOfSamples/2; s++ {
		samples = append(samples, 1)
	}

	for s := numOfSamples / 2; s < numOfSamples; s++ {
		samples = append(samples, -1)
	}

	return samples, nil
}

// Square8Bits returns the samples of a square wave with values ranging from
// 0 to 255. The duty cycle is 50%.
func Square8Bits(frequency float64, sampleRate uint) ([]uint8, error) {
	floatSamples, err := Square(frequency, sampleRate)
	if err != nil {
		return nil, err
	}

	return to8Bits(floatSamples), nil
}
