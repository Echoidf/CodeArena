package conf

import (
	"sync"
)

var viperMutex = &sync.Mutex{}

func SafeGetViperString(option string) string {
	viperMutex.Lock()
	defer viperMutex.Unlock()

	return V.GetString(option)
}
