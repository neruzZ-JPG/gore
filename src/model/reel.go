package model

import "gore/src/entity"

func SelectProcesses() []string {
	entity.Reel.Mutex.RLock()
	defer entity.Reel.Mutex.RUnlock()
	var res []string
	for key, value := range entity.Reel.ProcessMap {
		if value != entity.DELETE {
			res = append(res, key)
		}
	}
	return res
}

func RefreshFile() (err error) {
	err = entity.SaveFiles()
	return
}

func InsertProcess(name string) error {
	entity.Reel.Mutex.Lock()
	defer entity.Reel.Mutex.Unlock()
	entity.Reel.ProcessMap[name] = entity.NEW
	return nil
}

func DeleteProcess(name string) error {
	entity.Reel.Mutex.Lock()
	defer entity.Reel.Mutex.Unlock()
	entity.Reel.ProcessMap[name] = entity.DELETE
	return nil
}
