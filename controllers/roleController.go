package controllers

import (
	// "bytes"
	"fmt"
	"strconv"

	m "cstore/models"
)

type RoleController struct {
	CommonController
}

type PermissionTree struct {
	Fid       int64
	Pfid      int64
	Fname     string
	Ischecked bool
}

// 显示单个角色信息
func (this *RoleController) Show() {
	fmt.Println("this is user")
}

// 角色列表
func (this *RoleController) List() {
	if this.IsAjax() {
		sEcho := this.GetString("sEcho")
		// 表格显示在全部数据的开始位置
		iDiaplayStart := this.GetString("iDisplayStart")
		// 显示的条数
		iDisplayLength := this.GetString("iDisplayLength")
		iStart, _ := strconv.Atoi(iDiaplayStart)
		iLength, _ := strconv.Atoi(iDisplayLength)
		page := iStart / iLength

		_, count := m.GetAllRole()
		roles, _ := m.GetRoleList(int64(page+1), int64(iLength), "Id")
		fmt.Println("++++++++++++++")
		fmt.Println(roles)
		for _, role := range roles {
			switch role["Status"] {
			case int64(2):
				role["Statusname"] = "正常"
			case int64(1):
				role["Statusname"] = "禁止"

			default:
				role["Statusname"] = "不正常"
			}

		}

		data := make(map[string]interface{})
		data["aaData"] = &roles
		data["iTotalDisplayRecords"] = count
		data["iTotalRecords"] = iLength
		data["sEcho"] = sEcho
		this.Data["json"] = &data
		this.ServeJSON()
	} else {
		this.Data["ActionUrl"] = "localhost:8080"
		this.TplName = "role/index.html"
	}

}

// 添加角色
func (this *RoleController) Add() {
	isAction := this.GetString("isAction")
	if "0" == isAction {
		fmt.Println("添加角色开始")
		name := this.GetString("name")
		key := this.GetString("key")
		description := this.GetString("description")
		status, _ := this.GetInt64("status")

		_, role := m.GetRoleByName(name)
		if role.Id == 0 {
			role.Name = name
			role.Key = key
			role.Description = description
			role.Status = status

			_, err := m.AddRole(&role)
			if err != nil {
				fmt.Println("添加角色失败")
				fmt.Println(err)
				return
			}
		} else {
			fmt.Println("角色已经插入")
			this.Rsp(false, "角色已经插入")
		}

		this.Ctx.Redirect(302, "/cstore/role/add.html")
	} else {
		this.TplName = "role/add.html"
	}
}

//  删除角色
func (this *RoleController) Delete() {
	Id, _ := this.GetInt64("Id")
	_, err := m.DelRoleById(Id)
	if err != nil {
		fmt.Println("Err is")
		fmt.Println(err)
		this.Rsp(false, err.Error())
	} else {
		this.Rsp(true, "Delete Sucess")
		return
	}
}

// 更新角色
func (this *RoleController) Edit() {
	isAction := this.GetString("isAction")
	Id, _ := this.GetInt64("Id")
	Name := this.GetString("Name")
	Key := this.GetString("Key")
	Description := this.GetString("Description")
	Status, _ := this.GetInt64("Status")

	role := new(m.Role)
	role.Id = Id
	role.Name = Name
	role.Key = Key
	role.Description = Description
	role.Status = Status
	fmt.Println("++++++++++++++")
	fmt.Println(Status)
	if "0" == isAction {
		_, err := m.UpdateRole(role)
		if err != nil {
			this.Rsp(false, "更新出错："+err.Error())
		} else {
			this.Ctx.Redirect(302, "/cstore/role/list.html")
		}
	} else {
		this.Data["Resource"] = role
		this.TplName = "role/edit.html"
	}

}

// 给角色分配资源
func (this *RoleController) AllocationRes() {
	Id, _ := this.GetInt64("Id")
	fmt.Println("给角色分配资源权限")
	fmt.Printf("Id is %d\n", Id)

	res, _ := m.GetResourceByRoleId(Id)
	_, allRes := m.GetAllResource()
	tree := make([]PermissionTree, len(allRes))
	for k, allResOne := range allRes {
		var flag bool = false
		allId := allResOne["Id"].(int64)
		for _, resOne := range res {

			if resOne["Id"].(int64) == allId {
				tree[k].Fid = allId
				tree[k].Pfid = allResOne["Pid"].(int64)
				tree[k].Fname = allResOne["Name"].(string)
				tree[k].Ischecked = true
				flag = true
			}
		}
		if !flag {
			tree[k].Fid = allId
			tree[k].Pfid = allResOne["Pid"].(int64)
			tree[k].Fname = allResOne["Name"].(string)
		}
	}

	this.Data["roleId"] = Id

	this.Data["list"] = &tree
	fmt.Printf("String is %s", &tree)
	this.TplName = "role/permission.html"
}
