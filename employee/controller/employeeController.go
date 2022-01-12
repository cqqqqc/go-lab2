package controller

import (
	commonEntity "Lab2/common/entity"
	commonService "Lab2/common/service"
	"Lab2/employee/entity"
	"Lab2/employee/service"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"log"
	"net/http"
)

func CreateEmployee(c *gin.Context) {
	//定义一个User变量
	var employee entity.Employee
	//将调用后端的request请求中的body数据根据json格式解析到User结构变量中
	c.BindJSON(&employee)

	//客户端连接消息队列
	conn, err := amqp.Dial("amqp://guest:guest@127.0.0.1:5672")
	if err != nil {
		fmt.Println("dial")
		log.Fatal(err)
	}
	defer conn.Close()
	//获取通道，所有操作基本都是通道控制的
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	err = ch.ExchangeDeclare(
		"logs",   // name
		"fanout", // type
		false,    // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		log.Fatal(err)
	}
	//队列声明
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
	data := commonEntity.SimpleDemo{
		Name:       employee.Name,
		StuffNo:    employee.StuffNo,
		Department: employee.Department,
		Password:   "123456",
		Active:     false,
	}
	dataBytes, err := json.Marshal(data)
	if err != nil {
		commonService.ErrorHanding(err, "struct to json failed")
	}
	for i := 1; i <= 2; i++ {
		//发送消息
		err = ch.Publish(
			"logs",
			q.Name,
			false,
			false,
			amqp.Publishing{
				//DeliveryMode: amqp.Persistent,
				ContentType: "application/json",
				Body:        dataBytes, //消息的内容
			},
		)
	}

	if err != nil {
		log.Fatal(err)
	}

	//将被转换的user变量传给service层的CreateUser方法，进行User的新建
	employeeErr := service.CreateEmployee(&employee)
	//判断是否异常，无异常则返回包含200和更新数据的信息
	if employeeErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": employeeErr.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": employee,
		})
	}

}
func GetAllEmployee(c *gin.Context) {
	employeeList, err := service.GetAllEmployee()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": employeeList,
		})
	}
}
func UpdateEmployee(c *gin.Context) {
	//定义一个User变量
	var employee entity.Employee
	//将调用后端的request请求中的body数据根据json格式解析到User结构变量中
	c.BindJSON(&employee)
	//将被转换的user变量传给service层的CreateUser方法，进行User的新建
	err := service.UpdateEmployee(&employee)
	//判断是否异常，无异常则返回包含200和更新数据的信息
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": employee,
		})
	}
}
func DeleteEmployeeByStuffNo(c *gin.Context) {
	err := service.DeleteEmployeeByStuffNo(c.Param("stuffNo"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
		})
	}
}
func GetEmployeeByStuffNo(c *gin.Context) {
	user, err := service.GetEmployeeByStuffNo(c.Param("stuffNo"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": user,
		})
	}
}
