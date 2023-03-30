package file

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"errors"
	"os"
	"strings"

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
	// Read all the file
	in, err := os.Open(f.path)
	if err != nil {
		return nil, err
	}
	defer in.Close()

	// Read it into a string
	scan := bufio.NewScanner(in)
	scan.Split(bufio.ScanLines)

	// Loop through and return when we find the value
	for scan.Scan() {
		// Do something with the line
		line := scan.Text()
		if strings.HasPrefix(line, key) {
			// Decode the value
			return strings.Replace(line, key+"=", "", 1), nil
		}
	}

	return nil, errors.New("not found")
}

func (f *File) List() ([]qstore.Pair, error) {

	// Read all the file
	in, err := os.Open(f.path)
	if err != nil {
		return nil, err
	}
	defer in.Close()

	// Read it into a string
	scan := bufio.NewScanner(in)
	scan.Split(bufio.ScanLines)

	results := []qstore.Pair{}
	// Loop through and return when we find the value
	for scan.Scan() {
		// Do something with the line
		line := scan.Text()
		parts := strings.Split(line, "=")
		if len(parts) == 2 {
			results = append(results, qstore.Pair{
				Key:   parts[0],
				Value: parts[1],
			})
		}
	}

	return results, nil
}

func (f *File) Delete(key string) error {
	// Read all the file
	in, err := os.Open(f.path)
	if err != nil {
		return err
	}

	// Read it into a string
	scan := bufio.NewScanner(in)
	scan.Split(bufio.ScanLines)
	// Loop through the lines and append them to a string builder
	s := strings.Builder{}
	for scan.Scan() {
		// Do something with the line
		line := scan.Text()
		if !strings.HasPrefix(line, key) {
			s.WriteString(line)
			s.WriteString("\n")
		}
	}

	err = in.Close()
	if err != nil {
		return err
	}

	// Reopen the file with write permissions
	out, err := os.OpenFile(f.path, os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	out.Write([]byte(s.String()))
	return nil
}
