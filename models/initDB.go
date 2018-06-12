package models

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var odb orm.Ormer

/*
* 初始化数据库
*包括穿件数据库，表，以及插入部分数据
 */
func InitDB() {
	createdb()
	Connect()
	odb = orm.NewOrm()
	name := "default"
	force := true
	verbose := true
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Printf("init db is err ,this is %s ", err.Error())
	}

	insertUser()
	insertGroup()
	// insertRole()
	// insertNode()
	fmt.Println("database init is complete.\nPlease restart the application")
}

/**
* 创建数据库
 */
func createdb() {
	var sqlstring string
	dns, db_name := getConfig(0)
	sqlstring = fmt.Sprintf("CREATE DATABASE  if not exists `%s` CHARSET utf8 COLLATE utf8_general_ci", db_name)
	db, err := sql.Open("mysql", dns)
	if err != nil {
		panic(err.Error())
	}
	r, err := db.Exec(sqlstring)
	if err != nil {
		log.Printf("err is %s and r is %s", err.Error(), r)
	} else {
		log.Println("Database ", db_name, " created")
	}
	defer db.Close()

}
func Connect() {
	dns, _ := getConfig(1)
	fmt.Printf("数据库is %s", dns)
	err := orm.RegisterDataBase("default", "mysql", dns)
	if err != nil {
		fmt.Println("\n 数据库连接失败")
	} else {
		fmt.Println("\n 数据库连接sucess ")
	}
}

/*
* 获取配置
	flag ==1 表示 只链接
	==0 创建 加链接
*/
func getConfig(flag int) (string, string) {
	var dns string
	db_host := beego.AppConfig.String("db_host")
	db_port := beego.AppConfig.String("db_port")
	db_user := beego.AppConfig.String("db_user")
	db_pass := beego.AppConfig.String("db_pass")
	db_name := beego.AppConfig.String("db_name")
	if flag == 1 {
		fmt.Println("链接数据库")
		orm.RegisterDriver("mysql", orm.DRMySQL)
		dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=Local", db_user, db_pass, db_host, db_port, db_name)
	} else {
		fmt.Println("创建数据库")
		dns = fmt.Sprintf("%s:%s@tcp(%s:%s)?charset=utf8", db_user, db_pass, db_host, db_port)
	}

	return dns, db_name
}

func insertUser() {
	fmt.Println("insert user start .....")
	user := new(User)
	user.Uname = "admin"
	user.Nickname = "管理员"
	user.Pwd = "7f671d25bd47fc537bdf4e1f6723d899"
	user.Email = "qi19901212@163.com"
	user.Phone = "18510970061"
	user.Remark = "我是管理员，系统最大权限"
	user.Status = 2
	// o := orm.NewOrm()
	num, err := odb.Insert(user)
	if err != nil {
		fmt.Printf("num is %d and err is %s", num, err.Error())
	} else {
		fmt.Println("insert user end")
	}
}
func insertGroup() {
	fmt.Println("insert group  .....")
	g := new(Group)
	g.Name = "os"
	g.Title = "System"
	g.Sort = 1
	g.Status = 2
	num, err := odb.Insert(g)
	if err != nil {
		fmt.Printf("group insert ：num is %d and err is %s", num, err.Error())
	} else {
		fmt.Println("insert group end")
	}
}

/*func insertNodes() {
	fmt.Println("insert node ...")
	g := new(Group)
	g.Id = 1
	//nodes := make([20]Node)
	nodes := [24]Node{
		{Name: "rbac", Title: "RBAC", Remark: "", Level: 1, Pid: 0, Status: 2, Group: g},
		{Name: "node/index", Title: "Node", Remark: "", Level: 2, Pid: 1, Status: 2, Group: g},
		{Name: "index", Title: "node list", Remark: "", Level: 3, Pid: 2, Status: 2, Group: g},
		{Name: "AddAndEdit", Title: "add or edit", Remark: "", Level: 3, Pid: 2, Status: 2, Group: g},
		{Name: "DelNode", Title: "del node", Remark: "", Level: 3, Pid: 2, Status: 2, Group: g},
		{Name: "user/index", Title: "User", Remark: "", Level: 2, Pid: 1, Status: 2, Group: g},
		{Name: "Index", Title: "user list", Remark: "", Level: 3, Pid: 6, Status: 2, Group: g},
		{Name: "AddUser", Title: "add user", Remark: "", Level: 3, Pid: 6, Status: 2, Group: g},
		{Name: "UpdateUser", Title: "update user", Remark: "", Level: 3, Pid: 6, Status: 2, Group: g},
		{Name: "DelUser", Title: "del user", Remark: "", Level: 3, Pid: 6, Status: 2, Group: g},
		{Name: "group/index", Title: "Group", Remark: "", Level: 2, Pid: 1, Status: 2, Group: g},
		{Name: "index", Title: "group list", Remark: "", Level: 3, Pid: 11, Status: 2, Group: g},
		{Name: "AddGroup", Title: "add group", Remark: "", Level: 3, Pid: 11, Status: 2, Group: g},
		{Name: "UpdateGroup", Title: "update group", Remark: "", Level: 3, Pid: 11, Status: 2, Group: g},
		{Name: "DelGroup", Title: "del group", Remark: "", Level: 3, Pid: 11, Status: 2, Group: g},
		{Name: "role/index", Title: "Role", Remark: "", Level: 2, Pid: 1, Status: 2, Group: g},
		{Name: "index", Title: "role list", Remark: "", Level: 3, Pid: 16, Status: 2, Group: g},
		{Name: "AddAndEdit", Title: "add or edit", Remark: "", Level: 3, Pid: 16, Status: 2, Group: g},
		{Name: "DelRole", Title: "del role", Remark: "", Level: 3, Pid: 16, Status: 2, Group: g},
		{Name: "Getlist", Title: "get roles", Remark: "", Level: 3, Pid: 16, Status: 2, Group: g},
		{Name: "AccessToNode", Title: "show access", Remark: "", Level: 3, Pid: 16, Status: 2, Group: g},
		{Name: "AddAccess", Title: "add accsee", Remark: "", Level: 3, Pid: 16, Status: 2, Group: g},
		{Name: "RoleToUserList", Title: "show role to userlist", Remark: "", Level: 3, Pid: 16, Status: 2, Group: g},
		{Name: "AddRoleToUser", Title: "add role to user", Remark: "", Level: 3, Pid: 16, Status: 2, Group: g},
	}
	for _, v := range nodes {
		n := new(Node)
		n.Name = v.Name
		n.Title = v.Title
		n.Remark = v.Remark
		n.Level = v.Level
		n.Pid = v.Pid
		n.Status = v.Status
		n.Group = v.Group
		o.Insert(n)
	}
	fmt.Println("insert node end")
}
*/
