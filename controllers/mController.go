package controllers

import (
	m "cstore/models"
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	CommonController
}
type Tree struct {
	Id       int64
	Index    int    `json:"index"`
	Text     string `json:"id"`
	IconCls  string `json:"text"`
	Checked  string `json:"iconCls"`
	State    string `json:"checked"`
	Children []Tree `json:"state"`
	Url      string `json:"url"`
}

func (this *MainController) Index() {
	userinfo := this.GetSession("userinfo")
	if userinfo == nil {
		this.Ctx.Redirect(302, beego.AppConfig.String("auth_gateway"))
	}
	tree := this.GetTree()
	if this.IsAjax() {
		this.Data["json"] = &tree
		this.ServeJson()
	} else {
		groups := m.Groups()
		this.Data["userinfo"] = userinfo
		this.Data["groups"] = groups
		this.Data["tree"] = &tree
		this.TplNames = "index.tpl"
	}
}

func (this *MainController) Login() {
	isAjax := this.GetString("isAjax")
	if isAjax == "0" {
		uName := this.GetString("uname")
		pwd := this.GetString("password")
		user, err := m.CheckLogin(uName, pwd)
		if err == nil {
			this.SetSession("userinfo", user)
			accessList, _ := m.GetAccessList(user.Id)
			this.SetSession("accessList", accessList)
			this.Ctx.Redirect(302, "/cstore/mainFrame")
			return
		} else {
			this.Data["err"] = err
			this.TplNames = "err/error404.html"
			return
		}

	}
	userInfo := this.GetSession("userinfo")
	if userInfo != nil {
		this.Ctx.Redirect(302, "index.tpl")
	}
	this.TplNames = "login.html"
}
func (this *MainController) Logout() {
	this.DelSession("userinfo")
	this.Ctx.Redirect(302, "login.html")
}
func (this *MainController) ChangePwd() {
	fmt.Println("修改密码成功")
	userInfo := this.GetSession("userinfo")
	if userInfo == nil {
		this.Ctx.Redirect(302, beego.AppConfig.String("auth_gateway"))
	}
	oldPwd := this.GetString("oldpwd")
	newPwd := this.GetString("newpwd")
	repeatpassword := this.GetString("reppwd")
	if newPwd != repeatpassword {
		this.Rsp(false, "二次密码输入不一致")
	}
	user, err := m.CheckLogin(userInfo.(m.User).Uname, oldPwd)
	if err == nil {
		var u m.User
		u.Id = user.Id
		u.Pwd = newPwd
		id, err := m.UpdateUser(&u)
		if err == nil && id > 0 {
			this.Rsp(true, "密码修改成功")
			return
		} else {
			this.Rsp(false, err.Error())
			return
		}

	}
	this.Rsp(false, "密码错误")
}
func (this *MainController) MainFrame() {
	this.TplNames = "mainFrame.html"
}
func (this *MainController) Center() {
	this.TplNames = "background/center.html"
}

func (this *MainController) Top() {
	userInfo := this.GetSession("userinfo")

	if userInfo == nil {
		this.Ctx.Redirect(302, beego.AppConfig.String("auth_gateway"))
	} else {
		role := m.GetRoleByUserId(userInfo.(m.User).Id)
		roleName := role.Name
		this.Data["userinfo"] = userInfo
		this.Data["roleName"] = roleName
	}
	this.TplNames = "background/top.html"
}

func (this *MainController) Left() {
	userInfo := this.GetSession("userinfo")
	if userInfo == nil {
		this.Ctx.Redirect(302, beego.AppConfig.String("auth_gateway"))
	} else {
		role := m.GetRoleByUserId(userInfo.(m.User).Id)
		tree := this.GetResList(userInfo.(m.User).Uname, role.Id)

		this.Data["tree"] = &tree
		this.Data["ActionUrl"] = beego.AppConfig.String("ActionUrl")
	}
	this.TplNames = "background/left.html"
}
func (this *MainController) Tab() {
	this.TplNames = "background/tab.html"
}
