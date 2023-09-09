package conf

import (
	"CodeArena/consts"
	"CodeArena/utils"
	"os"
	"sync"
)

var viperMutex = &sync.Mutex{}

func SafeGetViperString(option string) string {
	viperMutex.Lock()
	defer viperMutex.Unlock()

	return V.GetString(option)
}

func GetLogFile() (logFile *os.File) {
	logPath := V.GetString(consts.LogPath)
	if utils.NotExistFile(logPath) {
		utils.CreateFile(logPath)
	}
	logFile, _ = os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	return
}
