package wavereader

import (
	"fmt"
	"io"
	"log"
)

// LoopReader is an io.Reader that reads n bytes from data
// before returning EOF. If n > len(data), data will be wrapped around.
type LoopReader struct {
	data       []byte
	dataIndex  uint
	virtualLen uint
}

func (r *LoopReader) Read(data []byte) (int, error) {
	for i := 0; i < len(data); i++ {
		if r.dataIndex < r.virtualLen {
			data[i] = r.data[r.dataIndex%uint(len(r.data))]
			r.dataIndex++
		} else {
			return i, io.EOF
		}
	}

	return len(data), nil
}

// NewLoopReader creates a LoopReader.
func NewLoopReader(data []byte, n uint) (*LoopReader, error) {
	log.Print("data:", data, "n:", n)
	if len(data) == 0 && n > 0 {
		return nil, fmt.Errorf("if n > 0, len(data) must be > 0")
	}
	reader := LoopReader{dataIndex: 0, virtualLen: n}
	reader.data = make([]byte, len(data))
	copy(reader.data, data)
	return &reader, nil
}