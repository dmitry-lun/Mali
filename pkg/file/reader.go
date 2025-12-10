package file

import "os"

type Reader interface {
	ReadFile(path string) ([]byte, error)
}

type FileReader struct{}

func NewFileReader() *FileReader {
	return &FileReader{}
}

func (r *FileReader) ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}
