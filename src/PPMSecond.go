package main

import (
	"fmt"
	"models"
	"github.com/go-martini/martini"
	"strconv"
)
 var taskList models.TaskList
func main() {
	taskList.InitTaskList()
	m := martini.Classic()
	SetRoute(m)
	m.Run()
}

func SetRoute(m *martini.ClassicMartini) {
	//主页
	m.Get("/", func() string {
			return "<b>欢迎使用任务跟踪管理系统 version 0.1<b>"
		})
	//所有任务
	m.Get("/task/all", func() string {

			return taskList.String()
		})
	
	//单个任务
	m.Get("/task/:id", func(params martini.Params) string {
			id,_:=strconv.Atoi( params["id"])

			if len(taskList.List)>id{
				one:=	taskList.List[id]
				return one.String()
			}else{
				return "<b>task not found.<b>"
			}
		})

	//用户登录
	m.Get("/hello/:name", func(params martini.Params) string {
			var u models.User
						u.Id=999
						u.Name=params["name"]
						fmt.Println(u)
			return "Hello " + params["name"]
		})
}