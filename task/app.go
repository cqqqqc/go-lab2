package main

import (
	"Lab2/task/dao"
	"Lab2/task/entity"
	"Lab2/task/route"
	"fmt"
	"github.com/go-ini/ini"
	"os"
)

func main() {
	//连接数据库
	dao.InitDB()
	//程序退出关闭数据库连接
	//defer dao.Close()
	//绑定模型
	dao.Db.AutoMigrate(&entity.Task{})
	httpPort, err := ini.Load("task/conf/app.init")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	runPort := httpPort.Section("server").Key("HttpPort").String()
	//注册路由
	r := route.SetRouter()
	//启动端口为8085的项目
	r.Run(":" + runPort) //读取端口号
}
