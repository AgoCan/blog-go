package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"blog-go/config"
	"blog-go/dal/rdb"
	"blog-go/models"
	"blog-go/routers"
	"blog-go/utils/session"
)

var (
	err error
)

func main() {
	session.Global.Close()
	session.Global = session.NewCookieManagerOptions(session.NewInMemStore(), &session.CookieMngrOptions{AllowHTTP: true})
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "c",
				Value:       "config/config.yaml",
				Usage:       "配置文件位置",
				Destination: &config.Opt.ConfigFile,
			},
		},
		Action: func(context *cli.Context) error {
			return nil
		},
	}
	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}

	// 初始化配置文件
	config.InitConfig(&config.Opt)

	// 连接数据库并在代码结束后关闭
	err = models.InitMysql()
	if err != nil {
		// 数据库连接失败，直接报错
		panic(err)
	}
	defer models.Close()

	// 链接redis
	err = rdb.InitRedis()
	if err != nil {
		// 数据库连接失败，直接报错
		panic(err)
	}
	defer rdb.Close()

	// 调用路由组
	router := routers.SetupRouter()

	err = router.Run(":9000")
	if err != nil {
		fmt.Println(err)
	}
}
