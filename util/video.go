package util

import (
	"log"
	"os"
	"path/filepath"
)

const VIDEO_DIR = "video"

func init() {

	err := os.Mkdir(VIDEO_DIR, 0750)
	log.Println(err.Error(), "视频文件夹初始化")
}

func GetVideoStoreDir() string {
	return filepath.Join("./", VIDEO_DIR)
}