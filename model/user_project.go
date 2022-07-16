package model

type UserProject struct {
	//gorm.Model
	UserId    int
	ProjectId int
	Role      string `gorm:"varchar(100);default:user"`
}
