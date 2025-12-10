package hash

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

type Hasher interface {
	HashMD5(data []byte) string
	HashSHA256(data []byte) string
	HashFileMD5(path string) (string, error)
	HashFileSHA256(path string) (string, error)
}

type FileHasher struct{}

func NewFileHasher() *FileHasher {
	return &FileHasher{}
}

func (h *FileHasher) HashMD5(data []byte) string {
	hash := md5.Sum(data)
	return fmt.Sprintf("%x", hash)
}

func (h *FileHasher) HashSHA256(data []byte) string {
	hash := sha256.Sum256(data)
	return fmt.Sprintf("%x", hash)
}

func (h *FileHasher) HashFileMD5(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

func (h *FileHasher) HashFileSHA256(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}
