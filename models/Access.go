package models

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	tool "cstore/src/tool"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

/*
*权限访问
 */

func CheckAccessAndRegisterRes() {
	var Check = func(ctx *context.Context) {
		user_auth_type, _ := strconv.Atoi(beego.AppConfig.String("user_auth_type"))
		auth_gateway := beego.AppConfig.String("auth_gateway")
		var accessList map[string]bool
		if user_auth_type > 0 {
			params := strings.Split(strings.ToLower(ctx.Request.RequestURI), "/")
			if CheckAccess(params) {
				// 获取session value
				uinfo := ctx.Input.Session("userinfo")
				if uinfo == nil {
					ctx.Redirect(302, auth_gateway)
				}
				adminUser := beego.AppConfig.String("admin_user")
				//直接通过认证
				if uinfo.(User).Uname == adminUser {
					return
				}
				if user_auth_type == 1 {
					sessions := ctx.Input.Session("accesss")
					if sessions != nil {
						accessList = sessions.(map[string]bool)
					}
				} else if user_auth_type == 2 {
					accessList, _ = GetAccessList(uinfo.(User).Id)
				}
				ret := AccessDecision(params, accessList)
				if !ret {
					ctx.Output.Json(&map[string]interface{}{"status": false, "info": "权限不足"}, true, false)
				}
			}
		}
	}
	beego.InsertFilter("/", beego.BeforeRouter, Check)
}

/*
*是否需要验证
 */
func CheckAccess(params []string) bool {
	if len(params) < 3 {
		return false
	}
	for _, nap := range strings.Split(beego.AppConfig.String("not_auth_package"), ",") {
		if params[1] == nap {
			return false
		}
	}
	return true
}

/*
*是否有权限 AccessDecision
 */
func AccessDecision(params []string, accessList map[string]bool) bool {
	if CheckAccess(params) {
		s := fmt.Sprintf("%s/%s/%s", params[1], params[2], params[3])
		if len(accessList) < 1 {
			return false
		}
		_, ok := accessList[s]
		if ok != false {
			return true
		}
	} else {
		return true
	}
	return false
}

type AccessNode struct {
	Id        int64
	Name      string
	Childrens []*AccessNode
}

/*
*访问权限列表
 */
func GetAccessList(uid int64) (map[string]bool, error) {
	list, err := AccessList(uid)
	if err != nil {
		return nil, err
	}
	alist := make([]*AccessNode, 0)
	for _, l := range list {
		if l["Pid"].(int64) == 0 && l["Level"].(int64) == 1 {
			aNode := new(AccessNode)
			aNode.Id = l["Id"].(int64)
			aNode.Name = l["Name"].(string)
			alist = append(alist, aNode)
		}
	}
	for _, l := range list {
		if l["Level"].(int64) == 2 {
			for _, an := range alist {
				if an.Id == l["Pid"].(int64) {
					aNode := new(AccessNode)
					aNode.Id = l["Id"].(int64)
					aNode.Name = l["Name"].(string)
					an.Childrens = append(an.Childrens, aNode)
				}
			}
		}
	}

	for _, l := range list {
		if l["Level"].(int64) == 3 {
			for _, an1 := range alist {
				if an1.Id == l["Pid"].(int64) {
					aNode := new(AccessNode)
					aNode.Id = l["Id"].(int64)
					aNode.Name = l["Name"].(string)
					an1.Childrens = append(an1.Childrens, aNode)
				}
			}
		}
	}
	accessList := make(map[string]bool)
	for _, v := range alist {
		for _, v1 := range v.Childrens {
			for _, v2 := range v1.Childrens {
				vname := strings.Split(v.Name, "/")
				v1name := strings.Split(v1.Name, "/")
				v2name := strings.Split(v2.Name, "/")
				str := fmt.Sprintf("%s/%s/%s", strings.ToLower(vname[0]), strings.ToLower(v1name[1]), strings.ToLower(v2name[0]))
				accessList[str] = true
			}
		}
	}
	return accessList, nil
}
func CheckLogin(uname string, pwd string) (user User, err error) {
	user = GetUserByUname(uname)
	if user.Id == 0 {
		return user, errors.New(" user is not exist")
	}
	password := tool.EncodeUserPwd(uname, pwd)
	fmt.Printf("\npassword is %s and Pwd is %s\n", password, user.Pwd)
	if user.Pwd != password {
		return user, errors.New("password is wrong ")
	}
	return user, nil
}
