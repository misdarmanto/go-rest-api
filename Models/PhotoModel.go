package Models

type PhotoModel struct {
	Id       uint   `gorm:"AUTO_INCREMENT;primaryKey;unique;not nul" json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photoUrl"`
	UserID   UserModel
}

func (b *PhotoModel) TableName() string {
	return "photo"
}
