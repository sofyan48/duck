package libs

import (
	"io"
	"os"

	"github.com/joho/godotenv"
)

// Check Error
// @e: error
func Check(e error) error {
	if e != nil {
		panic(e)
	}
	return e
}

// CheckFolder function check folder
// @path : string
// return error
func CheckFolder(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return err
	}
	return nil
}

// MakeDirs fucntion create directory
// @path : string
// return error
func MakeDirs(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

// FileRemove Remove Files
// @path : string
// return error
func FileRemove(path string) error {
	err := os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}

// CreateFile function create file
// @path : string
// return bool
func CreateFile(path string) bool {
	var _, err = os.Stat(path)

	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		Check(err)
		defer file.Close()
		return false
	}
	return true
}

// WriteFile func write local file
func WriteFile(path string, value string, perm os.FileMode) bool {
	var file, err = os.OpenFile(path, os.O_RDWR, perm)
	if Check(err) != nil {
		return false
	}
	defer file.Close()

	// write some text line-by-line to file
	_, err = file.WriteString(value)
	if Check(err) != nil {
		return false
	}
	// save changes
	err = file.Sync()
	if Check(err) != nil {
		return false
	}

	return true
}

// ReadFile function
func ReadFile(path string, perm os.FileMode) string {
	var file, err = os.OpenFile(path, os.O_RDWR, perm)
	if Check(err) != nil {
		return err.Error()
	}
	defer file.Close()
	var text = make([]byte, 1024)
	for {
		_, err = file.Read(text)
		if err == io.EOF {
			break
		}
		if err != nil && err != io.EOF {
			if Check(err) != nil {
				return err.Error()
			}
			break
		}
	}
	return string(text)
}

// DeleteFile Function
func DeleteFile(path string) bool {
	var err = os.Remove(path)
	if Check(err) != nil {
		return false
	}
	return true
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

// LoadEnvirontment load environment config
// @path : string
func LoadEnvirontment(path string) error {
	err := godotenv.Load()
	// err := godotenv.Load(path)
	Check(err)
	return err
}

// GetEnvirontment Get value from environtment
// @key : string
func GetEnvirontment(key string) string {
	var myEnv map[string]string
	myEnv, err := godotenv.Read()
	Check(err)
	return myEnv[key]
}

// GetAllEnvirontment Get value from environtment
// @key : string
func GetAllEnvirontment() map[string]string {
	godotenv.Load()
	var myEnv map[string]string
	myEnv, err := godotenv.Read()
	Check(err)
	return myEnv
}
