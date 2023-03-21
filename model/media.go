package model

import (
	"errors"
)

type Media struct {
	Model
	OriginName string
	Size int64
	Name string `gorm:"unique"`
	LocalUrl string
	UserID int
	User User
	Status int // -1 | 0 | 1 删除 ｜ 上传 ｜ 发布
}

func CreateMedia(mediaInfo *Media) error {
	if tx := DBInstance.Create(mediaInfo); tx.Error != nil {
		return tx.Error
	}
	return nil
}

func ProbMedia (mediaId string, userId string) error {
	media := &Media{}
	tx := DBInstance.First(media, mediaId)
	if tx.Error != nil {
		return errors.New(tx.Error.Error() + "当前视频id不存在")
	}
	tx = tx.Where("user_id", userId).First(&media)
	if tx.Error != nil {
		return errors.New(tx.Error.Error() + "当前用户无次视频")
	}
	return nil
}

func UpdateMedia(mediaId string, status int) (*Media, error) {
	media := &Media{
	}
	if tx := DBInstance.Model(media).Where("id = ?", mediaId).Update("status", status); tx.Error != nil {
		return nil, tx.Error
	}
	return media, nil
}