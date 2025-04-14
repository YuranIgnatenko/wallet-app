package utils

import (
	"os"
	"wallet-app/internal/logger"
)

var _prefix = "internal/services/queries/"

// used prefix for filepath !
func ReadQuery(filepath string) string {
	data, err := os.ReadFile(_prefix + filepath)
	if err != nil {
		logger.Log.Errorf("Error reading file query: %s", filepath)
		return ""
	}
	return string(data)

}
