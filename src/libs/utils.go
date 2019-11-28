package libs

import (
	"os"
)

// Check Error
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// CheckFolder function check folder
func CheckFolder(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return err
	}
	return nil
}

// MakeDirs fucntion create directory
func MakeDirs(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

// FileRemove Remove Files
func FileRemove(path string) error {
	err := os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}

// CheckEnvironment function check default env
// @path : string
// return bool, error
func CheckEnvironment(path string) (bool, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false, err
	}
	return true, nil
}
