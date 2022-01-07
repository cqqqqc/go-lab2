package main

import (
	"Lab2/employee/dao"
	"Lab2/employee/entity"
	"Lab2/employee/route"
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
	dao.Db.AutoMigrate(&entity.Employee{})
	httpPort, err := ini.Load("employee/conf/app.init")
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
