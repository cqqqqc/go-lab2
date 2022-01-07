package controller

import (
	"Lab2/employee/entity"
	"Lab2/employee/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateEmployee(c *gin.Context) {
	//定义一个User变量
	var employee entity.Employee
	//将调用后端的request请求中的body数据根据json格式解析到User结构变量中
	c.BindJSON(&employee)
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

	//var cUser *gin.Context
	//var user commonEntity.User
	//user.Name = employee.Name
	//user.StuffNo = employee.StuffNo
	//user.Department = employee.Department
	//user.Password = "123456"
	//user.Active = false
	//userErr := service.CreateUser(&user)
	//if userErr != nil {
	//	cUser.JSON(http.StatusBadRequest, gin.H{"error": userErr.Error()})
	//} else {
	//	cUser.JSON(http.StatusOK, gin.H{
	//		"code": 200,
	//		"msg":  "success",
	//		"data": user,
	//	})
	//}
	//
	//var cTask *gin.Context
	//var task commonEntity.Task
	//task.StuffNo = employee.StuffNo
	//task.TaskNo = string(rand.Int())
	//taskErr := service.CreateTask(&task)
	//if taskErr != nil {
	//	cTask.JSON(http.StatusBadRequest, gin.H{"error": taskErr.Error()})
	//} else {
	//	cTask.JSON(http.StatusOK, gin.H{
	//		"code": 200,
	//		"msg":  "success",
	//		"data": task,
	//	})
	//}
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
