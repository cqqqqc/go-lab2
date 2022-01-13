package main

import (
	"fmt"
	"github.com/go-ini/ini"
	"os"
	"user/dao"
	"user/entity"
	"user/rabbitMQ"
	"user/route"
)

//type Greeter struct{}
//
//func (g *Greeter) Greet(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
//	rsp.Msg = "Greet " + req.Name
//	return nil
//}

func main() {

	//连接数据库
	dao.InitDB()
	//程序退出关闭数据库连接
	//defer dao.Close()
	//绑定模型
	dao.Db.AutoMigrate(&entity.User{})
	httpPort, err := ini.Load("user/conf/app.init")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	runPort := httpPort.Section("server").Key("HttpPort").String()
	//注册路由
	r := route.SetRouter()

	rabbimq := rabbitMQ.NewRabbitMQRouting("name", "user")
	rabbimq.UserReceiveRouting()

	//启动端口为8085的项目
	r.Run(":" + runPort) //读取端口号

}
