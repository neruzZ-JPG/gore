package entity

import (
	"fmt"
	"gore/src/config"
	"io/ioutil"
	"os"
	"strings"
	"sync"
)

func InitData() {
	readProcessFile()
}

func readProcessFile() error {
	//read file or create first
	processMap, err := readFile(config.ReelPath)
	if err != nil && os.IsNotExist(err) {
		if _, erro := os.Create(config.ReelPath); erro != nil {
			return erro
		} else {
			processMap, _ = readFile(config.ReelPath)
		}
	}
	Reel = Reel_t{
		ProcessMap: processMap,
		Mutex:      sync.RWMutex{},
	}
	ConfigMap = make(map[string]*ConfigMap_t)
	for key := range processMap {
		ConfigMap[key] = nil
	}
	return nil
}

func ReadConfigFile(name string) error {
	configMap, err := readFile(config.ConfigDir + name)
	if err != nil {
		return err
	}
	cm := &ConfigMap_t{
		Configs: configMap,
		Mutex:   sync.RWMutex{},
	}
	ConfigMap[name] = cm
	return nil
}

func SaveFiles() error {
	Reel.Mutex.RLock()
	//save configs
	for key, value := range Reel.ProcessMap {
		if value == REFRESH {
			ConfigMap[key].Mutex.RLock()
			saveFile(ConfigMap[key].Configs, config.ConfigDir+key)
			ConfigMap[key].Mutex.RUnlock()
			Reel.ProcessMap[key] = NO_CHANGE
		}
		if value == NEW {
			//新增process记录
			//需要判断是否有config更改
			Reel.ProcessMap[key] = NO_CHANGE
			createFile(config.ConfigDir + key)
			configs, exist := ConfigMap[key]
			if exist {
				ConfigMap[key].Mutex.RLock()
				saveFile(configs.Configs, config.ConfigDir+key)
				ConfigMap[key].Mutex.RUnlock()
			}
		}
		if value == DELETE {
			if _, err := os.Open(config.ConfigDir + key); err != nil && os.IsNotExist(err) {
				delete(Reel.ProcessMap, key)
				continue
			}
			if err := deleteFile(config.ConfigDir + key); err != nil {
				return err
			}
			delete(Reel.ProcessMap, key)
		}
	}
	Reel.Mutex.RUnlock()
	// save reel
	Reel.Mutex.Lock()
	defer Reel.Mutex.Unlock()
	if err := saveFile(Reel.ProcessMap, config.ReelPath); err != nil {
		return err
	}
	return nil
}

//------------------------通用文件操作------------------
func readFile(path string) (map[string]string, error) {
	resMap := make(map[string]string, 0)
	file, err := os.OpenFile(path, os.O_RDONLY, 0400)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	if string(data) == "" {
		return resMap, nil
	}
	rows := strings.Split(string(data), "\n")
	for _, row := range rows {
		file := strings.Split(row, ":")
		resMap[file[0]] = file[1]
	}
	return resMap, nil
}

func saveFile(ssMap map[string]string, path string) error {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer file.Close()
	str := ""
	for k, v := range ssMap {
		str += fmt.Sprintf("%s:%s\n", k, v)
	}
	if str != "" {
		str = str[:len(str)-1]
	}
	if _, err = file.WriteString(str); err != nil {
		return err
	}
	return nil
}

func createFile(path string) error {
	if _, err := os.Open(path); err == nil {
		return fmt.Errorf("[file_createFile] file already exists")
	}
	if _, err := os.Create(path); err != nil {
		return fmt.Errorf("[file_createFile] error while creating file: %v", err)
	}
	return nil
}

func deleteFile(path string) error {
	if err := os.Remove(path); err != nil {
		return err
	}
	return nil
}
