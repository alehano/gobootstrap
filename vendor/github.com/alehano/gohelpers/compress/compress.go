package compress

import (
	"bytes"
	"compress/gzip"
	"errors"
	"github.com/golang/snappy"
	"io"
	"io/ioutil"
)

func Encode(data []byte, algorithm string) ([]byte, error) {
	if algorithm == "snappy" {
		return snappy.Encode(nil, data), nil
	} else if algorithm == "gz" {
		var buf bytes.Buffer
		w := gzip.NewWriter(&buf)
		_, err := w.Write(data)
		w.Close()
		if err != nil {
			return []byte{}, err
		}
		return buf.Bytes(), nil
	}
	return []byte{}, errors.New("Unsupported compress algorithm")
}

func EncodeFromReader(r io.Reader, algorithm string) (out []byte, err error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}
	return Encode(data, algorithm)
}

func Decode(data []byte, algorithm string) ([]byte, error) {
	if algorithm == "snappy" {
		return snappy.Decode(nil, data)
	} else if algorithm == "gz" {
		buf := bytes.NewBuffer(data)
		r, err := gzip.NewReader(buf)
		if err != nil {
			return []byte{}, err
		}
		return ioutil.ReadAll(r)
	}
	return []byte{}, errors.New("Unsupported compress algorithm")
}
