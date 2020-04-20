# wavegen    [![Build Status](https://travis-ci.org/rodrigo-dc/wavegen.svg?branch=master)](https://travis-ci.org/rodrigo-dc/wavegen)


Go package for discrete wave generation


Currently, `wavegen` is able to generate square and sine waves.

## Example

Simple example showing the generation of a sinusoid.

```go
package main

import (
	"fmt"
	"os"

	"github.com/rodrigo-dc/wavegen"
)

func main() {
	wave, err := wavegen.Sine(440, 44100)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Print tab-separated values, ready to be used by GNU Plot
	for i, s := range wave {
		fmt.Println(i, "\t", s)
	}
}
```

The output of the example above can be used as an input to GNU Plot.

```bash
./sine | gnuplot -e "set style data lines;plot '-'" --persist
```
![GNU Plot output](/examples/sine/sine.png?raw=true)