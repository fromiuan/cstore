package controllers

import (
	"fmt"
	"strconv"
	"strings"
	// "time"

	m "cstore/models"
)

type UserController struct {
	CommonController
}

func (this *UserController) List() {
	// iLength - iStart  + 1
	if this.IsAjax() {
		fmt.Println("ajax is true")
		sEcho := this.GetString("sEcho")
		iDisplayStart := this.GetString("iDisplayStart")
		iDisplayLength := this.GetString("iDisplayLength")
		iStart, _ := strconv.Atoi(iDisplayStart)
		iLength, _ := strconv.Atoi(iDisplayLength)
		page := iStart / iLength
		count, _ := m.Users()
		users, _ := m.GetUserList(int64(page+1), int64(iLength), "Id")
		for _, user := range users {
			switch user["Status"] {
			case int64(2):
				user["Statusname"] = "正常"
			case int64(1):
				user["Statusname"] = "禁止"
			default:
				user["Statusname"] = "不正常"
			}

			role := m.GetRoleByUserId(user["Id"].(int64))
			user["Rolename"] = role.Name
		}

		data := make(map[string]interface{})
		data["aaData"] = users
		data["iTotalDisplayRecords"] = count
		data["iTotalRecords"] = iLength
		data["sEcho"] = sEcho
		this.Data["json"] = &data
		this.ServeJSON()

	} else {
		this.Data["ActionUrl"] = "localhost:8080"
		this.TplName = "user/index.html"
	}
}

// 显示用户信息
func (this *UserController) Show() {

}

//就是代表更新
func (this *UserController) Edit() {

	isAction := this.GetString("isAction")

	Id, _ := this.GetInt64("Id")
	Uname := this.GetString("Uname")
	Pwd := this.GetString("Pwd")
	Nickname := this.GetString("Nickname")
	Email := this.GetString("Email")
	Phone := this.GetString("Phone")
	Status, _ := this.GetInt64("Status")
	Remark := this.GetString("Remark")
	// Logintime := this.GetString("Logintime")
	// Ctime := this.GetString("Ctime")

	user := new(m.User)
	user.Id = Id
	user.Uname = strings.TrimSpace(Uname)
	user.Pwd = strings.TrimSpace(Pwd)
	user.Nickname = strings.TrimSpace(Nickname)
	user.Email = Email
	user.Phone = Phone
	user.Status = Status
	user.Remark = Remark

	if "0" == isAction {
		_, err := m.UpdateUser(user)
		if err != nil {
			this.Rsp(false, "更新出现问题："+err.Error())
		} else {
			this.Ctx.Redirect(302, "/cstore/user/list.html")
		}
	} else {
		this.Data["Resource"] = user
		this.TplName = "user/edit.html"
	}

}

// 添加用户
func (this *UserController) Add() {
	isAction := this.GetString("isAction")
	if "0" == isAction {
		name := this.GetString("name")
		nickname := this.GetString("nickname")
		pwd := this.GetString("pwd")
		email := this.GetString("email")
		telphone := this.GetString("telphone")
		remark := this.GetString("remark")
		status, _ := this.GetInt64("status")

		// user := new(m.User)
		user := m.GetUserByUname(name)
		if user.Id == 0 {
			user.Uname = name
			user.Nickname = nickname
			user.Pwd = pwd
			user.Email = email
			user.Phone = telphone
			user.Remark = remark
			user.Status = status

			_, err := m.AddUser(&user)
			if err != nil {
				fmt.Println("插入数据库错误或者已经插入")
				fmt.Println(err)
				return
			}
		} else {
			fmt.Println("已经插入")
			this.Rsp(false, "已经插入")
		}

		this.Ctx.Redirect(302, "/cstore/user/add.html")
	} else {
		this.TplName = "user/add.html"
	}

}

// 删除用户
func (this *UserController) Delete() {
	Id, _ := this.GetInt64("Id")
	_, err := m.DelUserById(Id)
	if err != nil {
		this.Rsp(false, err.Error())
	} else {
		this.Rsp(true, "User is Delete Sucess")
	}

}

// 给用户分配角色
func (this *UserController) AllocationRole() {
	Id, _ := this.GetInt64("Id")
	fmt.Printf("给用户分配角色 Id is %d\n", Id)
	err, user := m.IsExitUser(Id)
	roleByUser := m.GetRoleByUserId(Id)

	roles, _ := m.GetAllRole()
	if err != nil {
		this.Rsp(false, "不存在这个用户")
	} else {
		this.Data["User"] = &user
		roleList := make([]m.Role, len(roles))
		for k, role := range roles {
			roleList[k].Id = role["Id"].(int64)
			roleList[k].Name = role["Name"].(string)
			roleList[k].Key = role["Key"].(string)
			roleList[k].Description = role["Description"].(string)
			roleList[k].Status = role["Status"].(int64)
			if roleList[k].Status == 2 {
				roleList[k].Isnormal = true
			} else {
				roleList[k].Isnormal = false
			}

		}
		this.Data["RoleList"] = &roleList
		this.Data["UserRole"] = roleByUser.Name
		this.TplName = "user/userRole.html"
	}
}

func (this *UserController) Allocation() {
	userId, _ := this.GetInt64("userId")
	roleId, _ := this.GetInt64("roleId")
	fmt.Printf("userId is %d and roleId is %d\n", userId, roleId)
	errId := m.DelUserRoleByUserId(userId)
	fmt.Println(errId)
	_, err := m.AddRoleUser(roleId, userId)
	// fmt.Println(err.Error())
	if err != nil {
		this.Rsp(false, "分配权限失败")
		fmt.Println("分配权限失败")
	} else {
		this.Rsp(true, "分配权限成功")
	}

}
