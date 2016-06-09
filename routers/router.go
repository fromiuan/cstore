package routers

import (
	"mime"
	"os"

	"cstore/controllers"
	"cstore/models"
	. "cstore/src/tool"
	"github.com/astaxie/beego"
)

func init() {

	initialize()
	router()
	beego.AddFuncMap("stringsToJson", StringtoJson)
}
func initialize() {

	mime.AddExtensionType(".css", "text/css")

	args := os.Args
	for _, v := range args {
		if v == "InitDB" {
			// 创建表 插入数据等等，只在执行./CStore InitDB 时候执行一次
			models.InitDB()
			os.Exit(0)
		}
	}
	// 链接数据库
	models.Connect()
}
func router() {
	beego.Router("/", &controllers.MainController{}, "*:Index")
	//登录路由
	beego.Router("/cstore/index", &controllers.MainController{}, "*:Index")
	beego.Router("/cstore/login", &controllers.MainController{}, "*:Login")
	beego.Router("/cstore/logout", &controllers.MainController{}, "*:Logout")
	beego.Router("/cstore/changepwd", &controllers.MainController{}, "*:ChangePwd")
	beego.Router("/cstore/mainFrame", &controllers.MainController{}, "*:MainFrame")
	//背景路由
	beego.Router("/cstore/background/top", &controllers.MainController{}, "*:Top")
	beego.Router("/cstore/background/center", &controllers.MainController{}, "*:Center")
	beego.Router("/cstore/background/left", &controllers.MainController{}, "*:Left")
	beego.Router("/cstore/background/tab", &controllers.MainController{}, "*:Tab")

	beego.Router("/cstore/user/list", &controllers.UserController{}, "*:List")
	beego.Router("/cstore/user/edit", &controllers.UserController{}, "*:Edit")
	beego.Router("/cstore/user/delete", &controllers.UserController{}, "*:Delete")
	beego.Router("/cstore/user/add", &controllers.UserController{}, "*:Add")
	beego.Router("/cstore/user/role", &controllers.UserController{}, "*:AllocationRole")
	beego.Router("/cstore/user/allocation", &controllers.UserController{}, "*:Allocation")

	beego.Router("/cstore/role/list", &controllers.RoleController{}, "*:List")
	beego.Router("/cstore/role/add", &controllers.RoleController{}, "*:Add")
	beego.Router("/cstore/role/delete", &controllers.RoleController{}, "*:Delete")
	beego.Router("/cstore/role/edit", &controllers.RoleController{}, "*:Edit")
	beego.Router("/cstore/role/resource", &controllers.RoleController{}, "*:AllocationRes")

	beego.Router("/cstore/resource/list", &controllers.ResController{}, "*:List")
	beego.Router("/cstore/resource/add", &controllers.ResController{}, "*:Add")
	beego.Router("/cstore/resource/delete", &controllers.ResController{}, "*:Delete")
	beego.Router("/cstore/resource/edit", &controllers.ResController{}, "*:Edit")
	beego.Router("/cstore/resource/saveRoleRescours", &controllers.ResController{}, "*:SaveRoleRescours")
	beego.Router("/cstore/resource/show", &controllers.ResController{}, "*:Show")
	//消息路由
	beego.Router("/cstore/message/list", &controllers.MesController{}, "*:List")
	beego.Router("/cstore/message/respond", &controllers.MesController{}, "*:Respond")
	beego.Router("/cstore/message/private", &controllers.MesController{}, "*:Private")
	beego.Router("/cstore/message/responddelect", &controllers.MesController{}, "*:Responddelect")
	beego.Router("/cstore/message/edit", &controllers.MesController{}, "*:Edit")
	beego.Router("/cstore/message/add", &controllers.MesController{}, "*:Add")
	beego.Router("/cstore/message/delete", &controllers.MesController{}, "*:Delete")
	beego.Router("/cstore/message/alldelete", &controllers.MesController{}, "*:Alldelete")
	beego.Router("/cstore/message/permissions", &controllers.MesController{}, "*:Permissions")
	beego.Router("/cstore/message/allpre", &controllers.MesController{},"*:Allpre")
    //学生管理
    //beego.Router("")
}
