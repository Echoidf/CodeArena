package utils

import (
	"github.com/spf13/viper"
	"sync"
)

var viperMutex = &sync.Mutex{}

func SafeGetString(option string) string {
	viperMutex.Lock()
	defer viperMutex.Unlock()

	return viper.GetString(option)
}
