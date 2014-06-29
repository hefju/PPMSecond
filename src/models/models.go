package models
import (
	"strconv"
	"bytes"
)

type User struct {
	Id int64
	Name string
}

type JTask struct {
	Id int64
	Title string
	Content string
}
func(t *JTask)String() string {
	return strconv.FormatInt(t.Id,10)+"\t"+ t.Title+",\t"+t.Content+"\n"
}

type  TaskList struct {
	List []JTask
}
func (t *TaskList)InitTaskList(){
	t0:= JTask{Id :1,Title:"t0",Content:"silce的用法"}
	t1:= JTask{Id :2,Title:"t1",Content:"struct的用法"}
	t2:=JTask{Id:3,Title:"t2",Content:"类型转换(int64 to string)"}
	t3:=JTask{Id:4,Title:"t3",Content:"字符串拼接(stringbuilder)"}
	t.List=append(t.List,t0,t1,t2,t3)
}
func (t *TaskList)String() string{
	if len(t.List)>0 {
		br:=bytes.NewBuffer(nil)
		br.WriteString("id"+"\t"+ "Title"+"\t"+"Content\n")
		for _,i:=range t.List{
			br.WriteString(strconv.FormatInt(i.Id,10)+"\t"+ i.Title+",\t"+i.Content+"\n")
	}
		return br.String()
	}else{
		return "没有内容"
	}
}


