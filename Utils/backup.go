package Utils

import (
	"io/ioutil"
	"os"
)

type IBackup interface {
	WriteBackup(value []byte) error
	ReadBackup() ([]byte, error)
	SetFileSection(string, string)
}

func (b Backup) WriteBackup(value []byte) error {

	_, err := os.Stat(b.Dir)
	if err != nil {
		os.Mkdir(b.Dir, 0644)
	}

	ioutil.WriteFile(b.Dir+b.File, value, 0644)
	return nil
}

func (b Backup) ReadBackup() ([]byte, error) {

	var file = b.Dir + b.File
	_, err := os.Stat(file)
	if err == nil {
		f, errfile := ioutil.ReadFile(file)

		if errfile != nil {
			return nil, errfile
		}

		return f, nil
	}
	return nil, err
}

func (b *Backup) SetFileSection(dir string, file string) {
	b.Dir = dir
	b.File = file
}
