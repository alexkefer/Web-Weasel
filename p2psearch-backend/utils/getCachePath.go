package utils

import (
	"github.com/alexkefer/p2psearch-backend/log"
	"os"
	"runtime"
)

func GetCachePath() (string, error) {
	home, err := os.UserHomeDir()

	if err != nil {
		log.Error("couldn't get users home directory")
	}

	switch runtime.GOOS {
	case "windows":
		return home + "\\p2pwebcache", err
	default:
		return home + "/.cache/p2pwebcache", err
	}
}
