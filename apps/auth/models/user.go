package models

type User struct {
	Base
	Name        string `json:"name"`
	IdentityDoc string `json:"identity_doc" gorm:"unique"`
	Email       string `json:"email" gorm:"unique"`
	Password    string `json:"password" gorm:"unique"`
}
