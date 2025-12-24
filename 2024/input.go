package main

import (
	"bufio"
	"encoding"
	"errors"
	"fmt"
	"io"
	"os"
	"sync"
)

type Input[K encoding.TextUnmarshaler] struct {
	file    *os.File
	scanner *bufio.Scanner
	sync.Mutex
}

func NewInput[K encoding.TextUnmarshaler](inputFile string) (*Input[K], error) {
	file, err := os.Open(inputFile)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to open input file: %s", inputFile))
	}

	scanner := bufio.NewScanner(file)

	return &Input[K]{
		file:    file,
		scanner: scanner,
	}, nil
}

func (i *Input[K]) Close() error {
	i.Lock()
	defer i.Unlock()
	err := i.file.Close()
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to close input file: %s", i.file.Name()))
	}

	return nil
}

func (i *Input[K]) Next() (K, error) {
	i.Lock()
	defer i.Unlock()

	i.scanner.Scan()
	val := i.scanner.Bytes()

	out := new(K)
	if len(val) == 0 {
		return out, io.EOF
	}

	if err := encoding.TextUnmarshaler(out).UnmarshalText(val); err != nil {
		return out, errors.New(fmt.Sprintf("Failed to unmarshal: %s", val))
	}

	return out, nil
}

func (i *Input[K]) All() ([]K, error) {
	var out []K
	for o, err := i.Next(); ; o, err = i.Next() {
		if err != nil {
			if errors.Is(err, io.EOF) {
				return out, nil
			}

			return out, err
		}

		out = append(out, o)
	}
}
