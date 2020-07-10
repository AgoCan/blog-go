package db

import (
	"blog-go/models"
	"fmt"
	"strings"
)

// InsertUser 插入用户信息
func InsertUser(user *models.User) (userID int64, err error) {
	if user == nil {
		err = fmt.Errorf("invalid user parameter")
		return
	}

	sqlstr := `INSERT INTO
						user(user_id, username, password, email, telephone, avatar, nick_name)
						values(?,?,?,?,?,?,?)`
	res, err := models.DB.Exec(sqlstr, user.UserID, user.Username, user.Password, user.Email,
		user.Telephone, user.Avatar, user.NickName)
	if err != nil {
		return
	}
	userID, err = res.LastInsertId()
	return
}

// GetUserList 获取用户列表
func GetUserList(pageNum, pageSize int) (userList []*models.User, err error) {
	if pageNum < 0 || pageSize < 0 {
		err = fmt.Errorf("invalid parameter, page_num:%d, page_size:%d", pageNum, pageSize)
		return
	}
	sqlstr := `select 
						id, created_at, updated_at, user_id,
						username, password, email, telephone,
						avatar,nick_name
					from 
						user 
					
					order by id desc
					limit ?, ?`
	err = models.DB.Select(&userList, sqlstr, pageNum, pageSize)
	return userList, err
}

// UpdateUser 更新
func UpdateUser(user *models.User) {
	var a []string
	a, err := models.ReflectTag(*user, "ID")
	if err != nil {
		return
	}
	s := strings.Join(a, ",")

	sqlstr := fmt.Sprintf("update user set %s where id = ?", s)
	fmt.Println(sqlstr)
	res, err := models.DB.Exec(sqlstr, user.ID)
	fmt.Println(res, err)
}
