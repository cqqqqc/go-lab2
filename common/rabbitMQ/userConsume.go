package rabbitMQ

import (
	commonEntity "Lab2/common/entity"
	"Lab2/user/entity"
	"Lab2/user/service"
	"encoding/json"
	"fmt"
	"reflect"
)

//路由模式Step：3、路由模式下消费代码
func (r *RabbitMQ) UserReceiveRouting() {
	//1、试探性创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		//交换机类型
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	r.failOnErr(err, "Failed t declare an exchange")

	//2、试探性创建队列，这里注意队列名称不要写
	q, err := r.channel.QueueDeclare(
		"", //随机生产队列名称
		true,
		false,
		true,
		false,
		nil,
	)
	r.failOnErr(err, "Failed to declare a queue")

	//3、绑定队列到exchange中
	err = r.channel.QueueBind(
		q.Name,
		//在Pub/Sub模式下，这里的key要为空
		r.Key,
		r.Exchange,
		false,
		nil,
	)

	//4、消费信息
	msgs, err := r.channel.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	//5、启动协程处理消息
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
			v.Ack(true)
		}
	}()
}
