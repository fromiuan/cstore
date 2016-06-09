package controllers

import (
	// "bytes"
	m "cstore/models"
	"fmt"
	"strconv"
	"time"
	"strings"

)

type MesController struct {
	CommonController
}

// 显示单个角色信息
func (this *MesController) Show() {
	// Uname := this.GetString("uname")
	fmt.Println("this is user")
}

//消息列表
func (this *MesController) List() {
	if this.IsAjax() {
		sEcho := this.GetString("sEcho")
		// 表格显示在全部数据的开始位置
		iDiaplayStart := this.GetString("iDisplayStart")
		// 显示的条数
		iDisplayLength := this.GetString("iDisplayLength")
		iStart, _ := strconv.Atoi(iDiaplayStart)
		iLength, _ := strconv.Atoi(iDisplayLength)
		page := iStart / iLength
		_, count := m.GetAllMassage()
		messages, _ := m.GetMessageList(int64(page+1), int64(iLength), "Id")

		for _, mess := range messages {
			mess["Addtime"] = mess["Addtime"].(time.Time).Format("2006-01-02 15:04:05")
			switch mess["Status"] {
			case int64(1):
				mess["Statuszhi"] = "隐私"
			case int64(2):
				mess["Statuszhi"] = "公开"

			default:
				mess["Statuszhi"] = "公开"
			}
		}

		data := make(map[string]interface{})
		data["aaData"] = &messages
		data["iTotalDisplayRecords"] = count
		data["iTotalRecords"] = iLength
		data["sEcho"] = sEcho
		this.Data["json"] = &data
		this.ServeJson()
	} else {
		this.Data["ActionUrl"] = "localhost:8080"
		this.TplNames = "message/index.html"
	}

}

//修改消息
func (this *MesController) Edit() {
	//消息开始
	Id, _ := this.GetInt64("Id")
	Name := this.GetString("Name")
	Status, _ := this.GetInt64("Status")
	Statusname := this.GetString("Statusname")
	message := new(m.Message)
	message.Id = Id
	message.Name = Name
	message.Status = Status
	message.Statusname = Statusname
	_, err := m.UpdateMessage(message)
	fmt.Println("-----------------------------------")
	fmt.Println(Statusname)
	fmt.Println(Name)
	if err != nil {
		this.Rsp(false, "更新出现问题："+err.Error())
	}
	this.Data["json"] = &message
	fmt.Println(message)
	this.ServeJson()

}

// // 添加消息
func (this *MesController) Add() {
		name := this.GetString("Name")
		statuszhi := this.GetString("Statuszhi")
		statusname := this.GetString("Statusname")
		status, _ := this.GetInt64("Status")
		resplay := this.GetString("Resplay")
		_, message := m.GetMessageByName(name)
		if message.Id == 0 {
			message.Name = name
			message.Statuszhi = statuszhi
			message.Statusname = statusname
			message.Status = status
			message.Resplay = resplay
			userinfo := this.GetSession("userinfo").(m.User)
			user := userinfo.Uname
			message.User = user
			this.Data["json"] = &message
			this.ServeJson()
			inserid, err := m.AddMessage(&message)
			if err != nil {
				fmt.Println("添加消息失败")
				fmt.Println(err)
				return
			}
			respond := new(m.Respond)
			if inserid > 0 {
				respond.Mesid = inserid
				respond.User = user
				_, err := m.AddRespond(respond)
				if err != nil {
					fmt.Println("添加消息失败")
					fmt.Println(err)
					return
				}
			}
			this.Rsp(true,"消息添加成功")
		} else {
			fmt.Println("消息已经插入")
			this.Rsp(false, "消息已经插入")
		}		
}

// //  修改权限
func (this *MesController) Permissions() {
	fmt.Println("permission is start ")
	Id, _ := this.GetInt64("Id")
	Name := this.GetString("Name")
	Status, _ := this.GetInt64("Status")
	statuszhi := this.GetString("statuszhi")
	fmt.Println("Id is ===", Id)
	message := new(m.Message)
	message.Id = Id
	message.Name = Name
	message.Status = Status
	message.Statuszhi = statuszhi
	_, err := m.UpdateMessage(message)
	if err != nil {
		this.Rsp(false, "更新出现问题："+err.Error())
	}
	this.Data["json"] = &message
	fmt.Println(message)
	this.ServeJson()

}
//批量修改权限
func (this *MesController) Allpre() {
	iDs := this.GetString("Id")
 	Preid := strings.Split(iDs, ",")
	fmt.Println("permission is start ")
	Status, _ := this.GetInt64("Status")
	for _,v:=range Preid{
		message := new(m.Message)
		id, _ := this.GetInt64("Id")	
		_, err := m.DelMessageById(id)
		id,err = strconv.ParseInt(v, 10, 64)
		message.Id=id
		message.Status = Status
		_, err = m.UpdateMessage(message)
		if err != nil {
		this.Rsp(false, "更新出现问题："+err.Error())
		this.Data["json"] = &message
		this.ServeJson()
		}
	}
	message := new(m.Message)
	this.Data["json"] = &message
	fmt.Println(message)
	this.ServeJson()
}
	
// //  删除消息
func (this *MesController) Delete() {
	Id, _ := this.GetInt64("Id")
	_, err := m.DelMessageById(Id)
	if err != nil {
		fmt.Println("Err is")
		fmt.Println(err)
		this.Rsp(false, err.Error())
	} else {
		this.Rsp(true, "Delete Sucess")
		return
	}
}

//批量删除
func (this *MesController) Alldelete() {

	iDs := this.GetString("Ids")
    Delid := strings.Split(iDs, ",")
    delelen := len(Delid)
    for i := 0; i < delelen-1; i++ {
    	fmt.Println(Delid)
    	id, _ := this.GetInt64("Id")
		for _,v:=range Delid{
			fmt.Println(v,"")
			a:=v
			fmt.Println(a)
			_, err := m.DelMessageById(id)
			id,err = strconv.ParseInt(a, 10, 64)
            fmt.Println(err)	
		}
		
        _, err := m.DelMessageById(id)
		if err != nil {
			fmt.Println("Err is")
			fmt.Println(err)
			this.Rsp(false, err.Error())
		} else {
			this.Rsp(true, "Delete Sucess")
			return
		}
    }
}

//  消息公共库
func (this *MesController) Respond() {
	if this.IsAjax() {
		sEcho := this.GetString("sEcho")
		// 表格显示在全部数据的开始位置
		iDiaplayStart := this.GetString("iDisplayStart")
		// 显示的条数
		iDisplayLength := this.GetString("iDisplayLength")
		iStart, _ := strconv.Atoi(iDiaplayStart)
		iLength, _ := strconv.Atoi(iDisplayLength)
		page := iStart / iLength
		m.Leftjoin(1, 10)
		_, count := m.GetAllMassage()
		messages := m.Leftjoin(int64(page+1), int64(iLength))
		for _, mess := range messages {

			switch mess["status"] {
			case "1":
				mess["statuszhi"] = "隐私"
			case "2":
				mess["statuszhi"] = "公开"
			}
		}
		data := make(map[string]interface{})
		data["aaData"] = &messages
		data["iTotalDisplayRecords"] = count
		data["iTotalRecords"] = iLength
		data["sEcho"] = sEcho
		this.Data["json"] = &data
		this.ServeJson()
	} else {
		this.Data["ActionUrl"] = "localhost:8080"
		this.TplNames = "message/respond.html"
	}

}

func (this *MesController) Private() {
	if this.IsAjax() {
		sEcho := this.GetString("sEcho")
		// 表格显示在全部数据的开始位置
		iDiaplayStart := this.GetString("iDisplayStart")
		// 显示的条数
		iDisplayLength := this.GetString("iDisplayLength")
		iStart, _ := strconv.Atoi(iDiaplayStart)
		iLength, _ := strconv.Atoi(iDisplayLength)
		page := iStart / iLength

		userinfo := this.GetSession("userinfo").(m.User)
		user := userinfo.Uname
		_, count := m.GetAllMassage()
		messages := m.Privatemes(int64(page+1), int64(iLength), user)
		for _, mess := range messages {

			switch mess["status"] {
			case "1":
				mess["statuszhi"] = "隐私"
			case "2":
				mess["statuszhi"] = "公开"
			}
		}
		data := make(map[string]interface{})
		data["aaData"] = &messages
		data["iTotalDisplayRecords"] = count
		data["iTotalRecords"] = iLength
		data["sEcho"] = sEcho
		this.Data["json"] = &data
		this.ServeJson()
	} else {
		this.Data["ActionUrl"] = "localhost:8080"
		this.TplNames = "message/private.html"
	}

}

//删除自动回复
func (this *MesController) Responddelect() {
	mesid, _ := this.GetInt64("mesid")

	_, err := m.DelMessageByMesid(mesid)
	if err != nil {
		fmt.Println("Err is")
		fmt.Println(err)
		this.Rsp(false, err.Error())
	} else {
		this.Rsp(true, "Delete Sucess")
		return
	}
}
