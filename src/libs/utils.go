package libs

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
	"github.com/sofyan48/dyno/src/libs/entity"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

// Check Error
// @e: error
func (util *Utils) Check(e error) error {
	if e != nil {
		util.LogFatal("Error : ", e)
	}
	return e
}

// LogInfo ...
func (util *Utils) LogInfo(word string, report interface{}) {
	log.Println(word, ":", report)
}

// LogFatal ...
func (util *Utils) LogFatal(word string, report interface{}) {
	log.Println(word, report)
}

// CheckFile function check folder
// @path : string
// return error
func (util *Utils) CheckFile(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		util.LogFatal("Error :", err)
		return false
	}
	return true
}

// MakeDirs fucntion create directory
// @path : string
// return error
func (util *Utils) MakeDirs(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

// FileRemove Remove Files
// @path : string
// return error
func (util *Utils) FileRemove(path string) error {
	err := os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}

// CreateFile function create file
// @path : string
// return bool
func (util *Utils) CreateFile(path string) bool {
	var _, err = os.Stat(path)

	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		util.Check(err)
		defer file.Close()
		return false
	}
	return true
}

// WriteFile func write local file
func (util *Utils) WriteFile(path string, value string, perm os.FileMode) bool {
	var file, err = os.OpenFile(path, os.O_RDWR, perm)
	if util.Check(err) != nil {
		return false
	}
	defer file.Close()

	// write some text line-by-line to file
	_, err = file.WriteString(value)
	if util.Check(err) != nil {
		return false
	}
	// save changes
	err = file.Sync()
	if util.Check(err) != nil {
		return false
	}

	return true
}

// ReadFile function
func (util *Utils) ReadFile(path string, perm os.FileMode) string {
	var file, err = os.OpenFile(path, os.O_RDWR, perm)
	if util.Check(err) != nil {
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
			if util.Check(err) != nil {
				return err.Error()
			}
			break
		}
	}
	return string(text)
}

// DeleteFile Function
func (util *Utils) DeleteFile(path string) bool {
	var err = os.Remove(path)
	if util.Check(err) != nil {
		return false
	}
	return true
}

// CheckEnvironment function check default env
// @path : string
// return bool, error
func (util *Utils) CheckEnvironment(path string) (bool, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false, err
	}
	return true, nil
}

// ReadHome function
// return string
func (util *Utils) ReadHome() string {
	usr, err := user.Current()
	util.Check(err)
	return usr.HomeDir
}

// LoadEnvirontment load environment config
// @path : string
func (util *Utils) LoadEnvirontment(path string) error {
	if path == "" {
		homeDir := util.ReadHome()
		err := godotenv.Load(homeDir + "/.dyno")
		util.Check(err)
		return err
	}
	err := godotenv.Load(path)
	util.Check(err)
	return err
}

// GetEnvirontment Get value from environtment
// @key : string
func (util *Utils) GetEnvirontment(key string) string {
	var myEnv map[string]string
	myEnv, err := godotenv.Read()
	util.Check(err)
	return myEnv[key]
}

// GetAllEnvirontment Get value from environtment
// @key : string
func (util *Utils) GetAllEnvirontment() map[string]string {
	godotenv.Load()
	var myEnv map[string]string
	myEnv, err := godotenv.Read()
	util.Check(err)
	return myEnv
}

// ServiceRegisterYML read YML File
// return map,error
func (util *Utils) ServiceRegisterYML(path string) (entity.ServiceRegisterYML, error) {
	taskRegister := entity.ServiceRegisterYML{}
	ymlFile, err := ioutil.ReadFile(path)
	if util.Check(err) != nil {
		return taskRegister, err
	}
	err = yaml.Unmarshal(ymlFile, &taskRegister)
	if util.Check(err) != nil {
		return taskRegister, err
	}
	return taskRegister, nil
}

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
func (util *Utils) GetCurrentPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	util.Check(err)
	return dir
}

// ConvertUnixTime ...
// @unixTime: int64
func (util *Utils) ConvertUnixTime(unixTime int64) time.Time {
	tm := time.Unix(unixTime, 0)
	return tm
}

// ParseJSON function conver json string to object
// @data: string
// return map[string]interface{}, error
func (util *Utils) ParseJSON(data string) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	err := json.Unmarshal([]byte(data), &result)
	if err != nil {
		util.Check(err)
		return nil, err
	}
	return result, nil
}

// CheckTemplateFile check template path
// @argsFile: string
func (util *Utils) CheckTemplateFile(path string) (string, error) {
	var templates string
	if path == "" {
		templates = util.GetCurrentPath() + "/dyno.yml"
	} else {
		templates = path
	}
	if !util.CheckFile(templates) {
		return "", cli.NewExitError("No Templates Parse", 1)
	}
	return templates, nil
}
