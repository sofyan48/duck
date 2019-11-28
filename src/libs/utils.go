package libs

import (
	"os"

	"github.com/joho/godotenv"
)

// Check Error
func Check(e error) error {
	if e != nil {
		panic(e)
	}
	return e
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
