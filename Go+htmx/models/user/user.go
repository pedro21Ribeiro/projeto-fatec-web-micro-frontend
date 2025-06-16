package userModels


type Tabler interface{
	TableName() string
}

type User struct {
	ID uint
	Name string `gorm:"column:nome"`
	Email string 
	Password string `gorm:"column:senha"`
}


func (User) TableName() string {
	return "usuarios"
}


type Users = []User

