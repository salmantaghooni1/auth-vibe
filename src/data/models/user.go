package models

type User struct {
	BaseModel
	Username     string `gorm:"size:30;type:string;not null;unique;"`
	FirstName    string `gorm:"size:20;type:string;null;"`
	LastName     string `gorm:"size:20;type:string;null;"`
	MobileNumber string `gorm:"size:11;type:string;null;unique;default:null;"`
	Email        string `gorm:"size:100;type:string;null;unique;default:null;"`
	Password     string `gorm:"size:100;type:string;not null;"`
	Picture      string `gorm:"size:100;type:string;"`
	Enabled      bool   `gorm:"default:true;"`
	UserRoles    *[]UserRole
}
