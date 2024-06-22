package model

type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password []byte `json:"-"`
}
