package server

import (
	"fmt"
	"math/rand"
	"os"
)

// issue: what if there is a problem during the write?
func SaveData1(path string, data []byte) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0664)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(data)
	return err
}

func SaveData2(path string, data []byte) error {
	tmp := fmt.Sprintf("%s.tmp.%d", path, rand.Int())
	f, err := os.OpenFile(tmp, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0664)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(data)
	if err != nil {
		os.Remove(tmp) // remove the file if something went wrong while writing
		return err
	}
	return os.Rename(tmp, path) // rename is atomic
}

// flush the file system before committing
func SaveData3(path string, data []byte) error {
	tmp := fmt.Sprintf("%s.tmp.%d", path, rand.Int())
	f, err := os.OpenFile(tmp, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0664)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(data)
	if err != nil {
		os.Remove(tmp)
		return err
	}
	err = f.Sync() // "fsync" or equivalent syscall
	if err != nil {
		os.Remove(tmp)
		return err
	}
	return os.Rename(tmp, path)
}

func LogCreate(path string) (*os.File, error) {
	return os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
}

func LogAppend(f *os.File, line string) error {
	buf := []byte(line)
	buf = append(buf, '\n')
	_, err := f.Write(buf)
	if err != nil {
		return err
	}
	return f.Sync()
}
