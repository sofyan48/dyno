package libs

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
)

// Check Error
// @e: error
func (lb *Libs) Check(e error) error {
	if e != nil {
		panic(e)
	}
	return e
}

// LogInfo ...
func (lb *Libs) LogInfo(word string, report interface{}) {
	log.Println(word, report)
}

// LogFatal ...
func (lb *Libs) LogFatal(word string, report interface{}) {
	log.Println(word, report)
}

// CheckFile function check folder
// @path : string
// return error
func (lb *Libs) CheckFile(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		lb.LogFatal("Error :", err)
		return false
	}
	return true
}

// MakeDirs fucntion create directory
// @path : string
// return error
func (lb *Libs) MakeDirs(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

// FileRemove Remove Files
// @path : string
// return error
func (lb *Libs) FileRemove(path string) error {
	err := os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}

// CreateFile function create file
// @path : string
// return bool
func (lb *Libs) CreateFile(path string) bool {
	var _, err = os.Stat(path)

	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		lb.Check(err)
		defer file.Close()
		return false
	}
	return true
}

// WriteFile func write local file
func (lb *Libs) WriteFile(path string, value string, perm os.FileMode) bool {
	var file, err = os.OpenFile(path, os.O_RDWR, perm)
	if lb.Check(err) != nil {
		return false
	}
	defer file.Close()

	// write some text line-by-line to file
	_, err = file.WriteString(value)
	if lb.Check(err) != nil {
		return false
	}
	// save changes
	err = file.Sync()
	if lb.Check(err) != nil {
		return false
	}

	return true
}

// ReadFile function
func (lb *Libs) ReadFile(path string, perm os.FileMode) string {
	var file, err = os.OpenFile(path, os.O_RDWR, perm)
	if lb.Check(err) != nil {
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
			if lb.Check(err) != nil {
				return err.Error()
			}
			break
		}
	}
	return string(text)
}

// DeleteFile Function
func (lb *Libs) DeleteFile(path string) bool {
	var err = os.Remove(path)
	if lb.Check(err) != nil {
		return false
	}
	return true
}

// CheckEnvironment function check default env
// @path : string
// return bool, error
func (lb *Libs) CheckEnvironment(path string) (bool, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false, err
	}
	return true, nil
}

// ReadHome function
// return string
func (lb *Libs) ReadHome() string {
	usr, err := user.Current()
	lb.Check(err)
	return usr.HomeDir
}

// LoadEnvirontment load environment config
// @path : string
func (lb *Libs) LoadEnvirontment(path string) error {
	if path == "" {
		homeDir := lb.ReadHome()
		err := godotenv.Load(homeDir + "/.duck")
		lb.Check(err)
		return err
	}
	err := godotenv.Load(path)
	lb.Check(err)
	return err
}

// GetEnvirontment Get value from environtment
// @key : string
func (lb *Libs) GetEnvirontment(key string) string {
	var myEnv map[string]string
	myEnv, err := godotenv.Read()
	lb.Check(err)
	return myEnv[key]
}

// GetAllEnvirontment Get value from environtment
// @key : string
func (lb *Libs) GetAllEnvirontment() map[string]string {
	godotenv.Load()
	var myEnv map[string]string
	myEnv, err := godotenv.Read()
	lb.Check(err)
	return myEnv
}

// // ReadYML read YML File
// // return map,error
// func (lb *Libs)ReadYML(path string) (scheme.RegisterTask, error) {
// 	var taskRegister scheme.RegisterTask
// 	ymlFile, err := ioutil.ReadFile(path)
// 	if Check(err) != nil {
// 		return taskRegister, err
// 	}
// 	err = yaml.Unmarshal(ymlFile, &taskRegister)
// 	if Check(err) != nil {
// 		return taskRegister, err
// 	}
// 	return taskRegister, nil
// }

// // ReadYMLSend read YML File
// // return map,error
// func ReadYMLSend(path string) (scheme.SendTask, error) {
// 	var taskSend scheme.SendTask
// 	ymlFile, err := ioutil.ReadFile(path)
// 	if Check(err) != nil {
// 		return taskSend, err
// 	}
// 	err = yaml.Unmarshal(ymlFile, &taskSend)
// 	if Check(err) != nil {
// 		return taskSend, err
// 	}
// 	return taskSend, nil
// }

// GetPCurrentPath get current path
// return string
func (lb *Libs) getPCurrentPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	lb.Check(err)
	return dir
}

// ConvertUnixTime ...
// @unixTime: int64
func (lb *Libs) ConvertUnixTime(unixTime int64) time.Time {
	tm := time.Unix(unixTime, 0)
	return tm
}

// ParseJSON function conver json string to object
// @data: string
// return map[string]interface{}, error
func (lb *Libs) ParseJSON(data string) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	err := json.Unmarshal([]byte(data), &result)
	if err != nil {
		lb.Check(err)
		return nil, err
	}
	return result, nil
}
