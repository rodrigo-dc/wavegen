package wavereader_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/rodrigo-dc/wavegen/wavereader"
)

var testData = []struct {
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
	for i, current := range testData {
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

func TestLoopReaderRestart(t *testing.T) {
	r, err := wavereader.NewLoopReader([]byte{0, 1, 2, 3}, 4)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	r.Restart()

	data := make([]byte, 2)
	_, err = r.Read(data)
	if bytes.Compare(data, []byte{0, 1}) != 0 {
		t.Error("expected:", []byte{0, 1}, "actual:", data)
	}

	r.Restart()

	data = make([]byte, 4)
	_, err = r.Read(data)
	if bytes.Compare(data, []byte{0, 1, 2, 3}) != 0 {
		t.Error("expected:", []byte{0, 1, 2, 3}, "actual:", data)
	}
}
