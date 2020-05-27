package models

//User is struct
type User struct {
	ID    int    `json:"id" gorm:"column:id;primary_key"`
	Name  string `json:"name" gorm:"column:name" sql:"not null;type:varchar(200)"`
	Email string `json:"email" gorm:"column:email" sql:"not null;type:varchar(200)"`
}

// User_DataTable users table type
type User_DataTable struct {
	User
	CommonColumn
}
