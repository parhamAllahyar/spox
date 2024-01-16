package file

import (
	"errors"
	"os"

)

type TempFile struct {
	Path  string
	Size  int64
	Title string
}

// size
func (f TempFile) Info() error {
	file, err := os.Stat(f.Path)
	if err != nil {
		return errors.New("")
	}
	f.Title = file.Name()
	f.Size = file.Size()
	return nil
}

func (f TempFile) LocalStorage() error {
	return errors.New("")
}

func (f TempFile) LocalRemover() error {
	return errors.New("")
}
