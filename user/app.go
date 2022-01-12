package main

import (
	commonEntity "Lab2/common/entity"
	"Lab2/user/dao"
	"Lab2/user/entity"
	"Lab2/user/route"
	"Lab2/user/service"
	"encoding/json"
	"fmt"
	"github.com/go-ini/ini"
	"github.com/streadway/amqp"
	"log"
	"os"
	"reflect"
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

	conn, err := amqp.Dial("amqp://guest:guest@127.0.0.1:5672")
	if err != nil {
		fmt.Println("dial")
		log.Fatal(err)
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	ch.ExchangeDeclare(
		"logs",
		"fanout",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	//此处队列声明中的参数必须和同名队列参数一致，否则将出错
	q, err := ch.QueueDeclare(
		"name",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	err = ch.QueueBind(
		q.Name,
		"",
		"logs",
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	err = ch.Qos(
		1,    //预取任务数量
		0,    //预取大小
		true, //全局设置
	)
	if err != nil {
		//无法设置Qos
		log.Fatal(err)
	}
	//消费者接收消息，msgs为只读通道
	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}
	//消费者会一直监听管道，启用协程
	go func() {
		var demo commonEntity.SimpleDemo
		for v := range msgs {
			fmt.Printf("body = %s\n", v.Body)
			fmt.Println(reflect.TypeOf(v.Body))
			err := json.Unmarshal(v.Body, &demo)
			//解析失败会报错，如json字符串格式不对，缺"号，缺}等。
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(demo.Password)
			var user entity.User
			user.Name = demo.Name
			user.StuffNo = demo.StuffNo
			user.Department = demo.Department
			user.Password = "123456"
			user.Active = false
			service.CreateUser(&user)
		}
	}()

	//启动端口为8085的项目
	r.Run(":" + runPort) //读取端口号

}
