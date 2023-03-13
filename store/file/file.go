package file

import (
	"bytes"
	"encoding/gob"
	"os"

	"github.com/andyantrim/qstore"
)

type File struct {
	path string
}

func NewFileWriter(path string, c chan qstore.Pair) (*File, error) {
	// Create the file if it doesn't exist
	_, err := os.OpenFile(path, os.O_CREATE, 0600)
	if err != nil {
		return nil, err
	}

	f := &File{
		path: path,
	}
	go f.listenAndWrite(c)
	return f, nil
}

func (f *File) listenAndWrite(c chan qstore.Pair) {
	for {
		p := <-c
		f.Set(p.Key, p.Value)
	}

}

func (f *File) Set(key string, value interface{}) error {
	of, err := os.OpenFile(f.path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer of.Close()

	// Convert everything to a byte slice
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	enc.Encode(value)
	if err != nil {
		return err
	}
	v := buf.Bytes()

	k := []byte(key + "=")
	b := make([]byte, len(k)+len(v)+1)

	// Copy the key, value and newline into the byte slice
	copy(b, k)
	copy(b[len(k):], v)
	copy(b[len(k)+len(v):], []byte("\n"))

	// Write to file
	if _, err := of.Write(b); err != nil {
		return err
	}

	// Return the value
	return nil

}

func (f *File) Get(key string) (interface{}, error) {
	// read from file
	return "", nil
}
