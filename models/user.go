package models

// User 用户表
type User struct {
	baseModel
	UserID    int64  `db:"user_id"`
	Username  string `db:"username"  json:"username"`
	Password  string `db:"password"  json:"password"`
	Email     string `db:"email"     json:"email"`
	Telephone int    `db:"telephone" json:"telephone"`
	Avatar    string `db:"avatar"    json:"avatar"`
	NickName  string `db:"nick_name" json:"nick_name"`
}
