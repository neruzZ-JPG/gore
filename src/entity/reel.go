package entity

import "sync"

const (
	NO_CHANGE = "0"
	REFRESH   = "1"
	NEW       = "2"
	DELETE    = "-1"
)

type Reel_t struct {
	ProcessMap map[string]string
	Mutex      sync.RWMutex
}

var Reel Reel_t

var ConfigMap map[string]*ConfigMap_t
