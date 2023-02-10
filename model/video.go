package model

type Video struct {
	Model
	OriginName string
	Size int64
	Name string `gorm:"unique"`
	LocalUrl string
	UserID int
	User User
}

func CreateVideo(videoInfo *Video) error {
	if tx := DBInstance.Create(videoInfo); tx.Error != nil {
		return tx.Error
	}
	return nil
}