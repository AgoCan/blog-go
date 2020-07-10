package db

import (
	"blog-go/models"
	"fmt"
	"net/url"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func init() {
	timezone := "'Asia/Shanghai'"
	dns := "root:123456@tcp(127.0.0.1:3307)/blog_go?charset=utf8mb4&parseTime=True&loc=Local&time_zone=" + url.QueryEscape(timezone)
	var err error
	models.DB, err = sqlx.Open("mysql", dns)
	// fmt.Println(dns)
	err = models.DB.Ping()
	if err != nil {
		panic(err)
	}
	// INSERT INTO `blog_go`.`category`(`id`, `category_name`, `category_no`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, 'css/html', 1, '2018-08-12 10:55:45', '2018-08-12 10:59:00', NULL);
}

func TestInsertUser(t *testing.T) {
	// dns := "root:root1234@tcp(127.0.0.1:3307)/blog_go?charset=utf8mb4&parseTime=True&loc=Local"
	// var err error
	// models.DB, err = sqlx.Open("mysql", dns)
	// fmt.Println(err)
	var user models.User
	user.Username = "hank"
	user.Avatar = "/img/default.png"
	user.NickName = "hank997"
	user.Password = "123456"
	user.Telephone = 123456
	user.Email = "hank@hank.cn"
	user.UserID = 123333
	i, e := InsertUser(&user)
	fmt.Println(i, e)
}

func TestGetUserList(t *testing.T) {
	userList, err := GetUserList(0, 2)
	if err != nil {
		t.Error("err", err)
		return
	}
	for _, i := range userList {
		t.Log(i.ID)
	}
	t.Log("user", len(userList))
}

func TestUpdateUser(t *testing.T) {
	var user models.User
	user.ID = 1
	//user.NickName = "ago"
	UpdateUser(&user)
}
