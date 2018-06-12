package controllers

import (
	"fmt"
	"strconv"
	"strings"

	m "cstore/models"
	// "github.com/astaxie/beego"
)

type ResController struct {
	CommonController
}

// 资源列表
func (this *ResController) List() {
	if this.IsAjax() {
		sEcho := this.GetString("sEcho")
		iDisplayStart := this.GetString("iDisplayStart")
		iDisplayLength := this.GetString("iDisplayLength")
		iStart, _ := strconv.Atoi(iDisplayStart)
		iLength, _ := strconv.Atoi(iDisplayLength)
		page := iStart / iLength
		count, _ := m.GetAllResource()
		res, _ := m.GetResources(int64(page+1), int64(iLength), "Level")

		for _, resource := range res {
			switch resource["Type"] {
			case "0":
				resource["Typename"] = "目录"
			case "1":
				resource["Typename"] = "菜单"
			case "2":
				resource["Typename"] = "按钮"
			}
		}

		data := make(map[string]interface{})
		data["aaData"] = res
		data["iTotalDisplayRecords"] = count
		data["iTotalRecords"] = iLength
		data["sEcho"] = sEcho
		this.Data["json"] = &data
		this.ServeJSON()

	} else {
		this.Data["ActionUrl"] = "localhost:8080"
		this.TplName = "res/index.html"
	}
}

// 显示单个资源信息
func (this *ResController) Show() {

}

// 添加资源
func (this *ResController) Add() {
	isAction := this.GetString("isAction")
	if "0" == isAction {
		name := this.GetString("name")
		pid, _ := this.GetInt64("pid")
		key := this.GetString("key")
		url := this.GetString("url")
		level, _ := this.GetInt64("level")
		decription := this.GetString("description")
		restype := this.GetString("restype")

		_, resource := m.GetResourceByName(name)
		if resource.Id == 0 {
			resource.Name = name
			resource.Pid = pid
			resource.Key = key
			resource.Url = url
			resource.Level = level
			resource.Description = decription
			resource.Type = restype
			group := new(m.Group)
			group.Id = 1
			resource.Group = group

			_, err := m.AddResource(&resource)
			if err != nil {
				fmt.Println("添加资源失败")
				fmt.Println(err)
				return
			}
		} else {
			fmt.Println("资源已经插入")
			this.Rsp(false, "资源已经插入")
		}

		this.Ctx.Redirect(302, "/cstore/res/add.html")
	} else {

		this.TplName = "res/add.html"
	}

}

// 更新资源
func (this *ResController) Edit() {
	isAction := this.GetString("isAction")

	fmt.Println("Edit is start ")

	Id, _ := this.GetInt64("Id")
	Name := this.GetString("Name")
	Pid, _ := this.GetInt64("Pid")
	Key := this.GetString("Key")
	Type := this.GetString("Type")
	Url := this.GetString("Url")
	Level, _ := this.GetInt64("Level")
	Description := this.GetString("Description")

	fmt.Println("Id is ===", Id)
	fmt.Println("============================================")
	fmt.Println("Name is ===", Name)
	// fmt.Println("Url is ===", Url)
	// fmt.Println("Description is ===", Description)

	resource := new(m.Resource)
	resource.Id = Id
	resource.Name = Name
	resource.Pid = Pid
	resource.Key = Key
	resource.Url = Url
	resource.Level = Level
	resource.Description = Description
	resource.Type = Type
	if "0" == isAction {
		fmt.Println("Pid is ===", Pid)
		_, err := m.UpdateResource(resource)
		if err != nil {
			this.Rsp(false, "更新出现问题："+err.Error())
		} else {
			this.Ctx.Redirect(302, "/cstore/resource/list")
		}
	} else {
		_, pRes := m.GetParentResource(0)
		this.Data["pRes"] = pRes
		this.Data["Resource"] = resource
		this.TplName = "res/edit.html"
	}

}

// 删除资源
func (this *ResController) Delete() {
	Id := this.GetString("Id")
	resId, _ := strconv.Atoi(Id)
	status, err := m.DelResourceById(int64(resId))
	fmt.Printf("Id is %s and status is %d\n", Id, status)
	if err == nil && status > 0 {
		this.Rsp(true, "Success")
		return
	} else {
		this.Rsp(false, err.Error())
		return
	}

}

func (this *ResController) SaveRoleRescours() {
	roleId, _ := this.GetInt64("roleId")
	fid := this.GetString("rescId")
	resList := strings.Split(fid, ",")
	resList = resList[0 : len(resList)-1]
	s := make([]string, 1)
	var isCotaions bool
	fmt.Printf("resList Id is %s\n", resList)
	for _, v := range resList {

		for _, resId := range s {
			if strings.EqualFold(v, resId) {
				isCotaions = true
			}
		}
		if !isCotaions {
			if s[0] == "" {
				s[0] = v
			} else {
				s = append(s, v)
			}

		} else {
			isCotaions = false
		}
	}
	err := m.DelRoleRescoursByRoleId(roleId)
	m.AddRoleRescours(roleId, s)

	if err != nil {
		this.Rsp(false, "fail")
	} else {
		this.Rsp(true, "成功")
	}

}
