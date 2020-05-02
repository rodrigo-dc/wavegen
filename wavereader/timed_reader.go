package wavereader

import (
	"time"
)

// TimedReader is an io.Reader that reads the number of bytes that corresponds
// to a certain duration. It uses LoopReader with an n that depends on the
// duration and sampleRate.
type TimedReader struct {
	loopReader *LoopReader
}

// NewTimedReader creates a new TimedReader.
func NewTimedReader(data []byte, sampleRate uint, d time.Duration) (*TimedReader, error) {
	sampleCount := float64(sampleRate) * d.Seconds()
	r, err := NewLoopReader(data, uint(sampleCount))
	if err != nil {
		return nil, err
	}

	return &TimedReader{loopReader: r}, nil
}

func (r *TimedReader) Read(data []byte) (int, error) {
	return r.loopReader.Read(data)
}

// Restart resets the TimedReader to its initial state.
// After calling Restart, n bytes can be read from the TimedReader again.
func (r *TimedReader) Restart() {
	r.loopReader.Restart()
}
