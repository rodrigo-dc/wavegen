package wavereader_test

import (
	"bytes"
	"io"
	"testing"
	"time"

	"github.com/rodrigo-dc/wavegen/wavereader"
)

type waveReader interface {
	io.Reader
	Restart()
}

var testLoopReaderData = []struct {
	data     []byte
	n        uint
	expected []byte
}{
	{
		[]byte{11, 22, 33, 44, 55},
		5,
		[]byte{11, 22, 33, 44, 55},
	},
	{
		[]byte{11, 22, 33, 44, 55},
		1,
		[]byte{11},
	},
	{
		[]byte{11, 22, 33, 44, 55},
		1,
		[]byte{11},
	},
	{
		[]byte{11, 22, 33, 44, 55},
		10,
		[]byte{11, 22, 33, 44, 55, 11, 22, 33, 44, 55},
	},
	{
		[]byte{11, 22, 33, 44, 55},
		0,
		[]byte{},
	},
}

func TestLoopReader(t *testing.T) {
	for i, current := range testLoopReaderData {
		r, err := wavereader.NewLoopReader(current.data, current.n)
		if err != nil {
			t.Errorf("test: %d, unexpected error in NewLoopReader: %s", i,
				err.Error())
		}

		output := make([]byte, current.n)
		n, err := r.Read(output)
		if err != nil {
			t.Errorf("test: %d, unexpected error in Read: %s", i, err.Error())
		}

		if uint(n) != current.n {
			t.Errorf("test: %d, expected n: %d, actual n: %d", i, current.n, n)
		}
		if !bytes.Equal(output, current.expected) {
			t.Error("test:", i, "expected:", current.expected, "actual:", output)
		}
	}
}

func TestLoopReaderErrors(t *testing.T) {
	var r *wavereader.LoopReader
	var err error
	// empty slice, n > 0
	r, err = wavereader.NewLoopReader([]byte{}, 5)
	if err == nil {
		t.Error("An error was expected in this case")
	}
	if r != nil {
		t.Error("Non-nil reader returned in an error condition")
	}
	// empty slice, n == 0
	r, err = wavereader.NewLoopReader([]byte{}, 0)
	if err != nil {
		t.Errorf("unexpected error in NewLoopReader: %s", err.Error())
	}
	if r == nil {
		t.Error("unexpected nil reader")
	}
	var n int
	n, err = r.Read(make([]byte, 10))
	if n != 0 {
		t.Error("expected zero read size, got", n)
	}
	if err != io.EOF {
		t.Errorf("expected EOF, got %v", err)
	}
	// len(slice) > 0, n == 0
	r, err = wavereader.NewLoopReader([]byte{1, 2, 3}, 0)
	if err != nil {
		t.Errorf("unexpected error in NewLoopReader: %s", err.Error())
	}
	if r == nil {
		t.Error("unexpected nil reader")
	}
	n, err = r.Read(make([]byte, 10))
	if n != 0 {
		t.Error("expected zero read size, got", n)
	}
	if err != io.EOF {
		t.Errorf("expected EOF, got %v", err)
	}
}

func TestTimedReaderErrors(t *testing.T) {
	var r *wavereader.TimedReader
	var err error
	// empty slice, expected number of samples > 0
	r, err = wavereader.NewTimedReader([]byte{}, 5, 1*time.Second)
	if err == nil {
		t.Error("An error was expected in this case")
	}
	if r != nil {
		t.Error("Non-nil reader returned in an error condition")
	}
	// empty slice, expected number of samples == 0
	r, err = wavereader.NewTimedReader([]byte{}, 0, 0*time.Second)
	if err != nil {
		t.Errorf("unexpected error in NewTimedReader: %s", err.Error())
	}
	if r == nil {
		t.Error("unexpected nil reader")
	}
	var n int
	n, err = r.Read(make([]byte, 10))
	if n != 0 {
		t.Error("expected zero read size, got", n)
	}
	if err != io.EOF {
		t.Errorf("expected EOF, got %v", err)
	}
	// len(slice) > 0, expected number of samples == 0
	r, err = wavereader.NewTimedReader([]byte{1, 2, 3}, 5, 0*time.Second)
	if err != nil {
		t.Errorf("unexpected error in NewTimedReader: %s", err.Error())
	}
	if r == nil {
		t.Error("unexpected nil reader")
	}
	n, err = r.Read(make([]byte, 10))
	if n != 0 {
		t.Error("expected zero read size, got", n)
	}
	if err != io.EOF {
		t.Errorf("expected EOF, got %v", err)
	}
}

func testRestart(t *testing.T, r waveReader) {
	r.Restart()

	data := make([]byte, 2)
	r.Read(data)
	if bytes.Compare(data, []byte{0, 1}) != 0 {
		t.Error("expected:", []byte{0, 1}, "actual:", data)
	}

	r.Restart()

	data = make([]byte, 4)
	r.Read(data)
	if bytes.Compare(data, []byte{0, 1, 2, 3}) != 0 {
		t.Error("expected:", []byte{0, 1, 2, 3}, "actual:", data)
	}
}

func TestLoopReaderRestart(t *testing.T) {
	r, err := wavereader.NewLoopReader([]byte{0, 1, 2, 3}, 4)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	testRestart(t, r)
}

func TestTimedReaderRestart(t *testing.T) {
	r, err := wavereader.NewTimedReader([]byte{0, 1, 2, 3}, 4, 1*time.Second)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	testRestart(t, r)
}

var testTimedReaderData = []struct {
	data       []byte
	sampleRate uint
	duration   time.Duration
	expected   []byte
}{
	{
		[]byte{128, 255, 128, 0},
		4,
		time.Second * 1,
		[]byte{128, 255, 128, 0},
	},
	{
		[]byte{128, 255, 128, 0},
		4,
		time.Second * 2,
		[]byte{128, 255, 128, 0, 128, 255, 128, 0},
	},
	{
		[]byte{128, 255, 128, 0},
		2,
		time.Second * 2,
		[]byte{128, 255, 128, 0},
	},
}

func TestTimedReader(t *testing.T) {
	for i, current := range testTimedReaderData {
		r, err := wavereader.NewTimedReader(current.data, current.sampleRate, current.duration)
		if err != nil {
			t.Errorf("test: %d, unexpected error in NewTimedReader: %s", i,
				err.Error())
		}

		output := make([]byte, len(current.expected))
		n, err := r.Read(output)
		if err != nil {
			t.Errorf("test: %d, unexpected error in Read: %s", i, err.Error())
		}

		if n != len(current.expected) {
			t.Errorf("test: %d, expected n: %d, actual n: %d", i, len(current.expected), n)
		}
		if !bytes.Equal(output, current.expected) {
			t.Error("test:", i, "expected:", current.expected, "actual:", output)
		}
	}
}
