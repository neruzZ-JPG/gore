package entity

import "sync"

type ConfigMap_t struct {
	Configs map[string]string
	Mutex   sync.RWMutex
}
