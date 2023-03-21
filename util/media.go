package util

import (
	"github.com/spf13/viper"
	"path/filepath"
)



func GetMediaStoreDir() string {
	return filepath.Join("./", viper.GetString("media.dir"))
}