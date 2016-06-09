package models

import (
	"errors"
	"fmt"
	"log"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

type Role struct {
	Id          int64
	Name        string `orm:"size(64)" form:"Name" valid:"Required" `
	Key         string `orm:"size(64)" form:"Key" valid:"Required"`
	Description string `orm:"size(200)" form:"Description" valid:"MaxSize(200)"`
	//  Status  2 正常  1 禁用
	Status     int64       `orm:"default(2)" form:"Status" valid:"Range(1,2)"`
	Statusname string      `orm:"-" form:"Statusname" `
	Isnormal   bool        `orm:"-" form:"Isnormal" valid:"Required" `
	User       []*User     `orm:"reverse(many)"`
	Resource   []*Resource `orm:"reverse(many)"`
}

func (r *Role) TableName() string {
	return beego.AppConfig.String("role_table")
}

/*
*注册model
 */
func init() {
	orm.RegisterModel(new(Role))
}

/*
*检查role表
 */
func checkRole(r *Role) (err error) {
	valid := validation.Validation{}
	b, _ := valid.Valid(&r)
	if !b {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			return errors.New(err.Message)
		}
	}
	return nil
}

func GetRoleList(page int64, page_size int64, sort string) (roles []orm.Params, count int64) {
	o := orm.NewOrm()
	role := new(Role)
	qs := o.QueryTable(role)
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}
	qs.Limit(page_size, offset).OrderBy(sort).Values(&roles)
	count, _ = qs.Count()
	return roles, count
}
func AddRole(r *Role) (int64, error) {
	if err := checkRole(r); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	role := new(Role)
	role.Name = r.Name
	role.Key = r.Key
	role.Description = r.Description
	role.Status = r.Status

	id, err := o.Insert(role)
	return id, err
}
func UpdateRole(r *Role) (int64, error) {
	if err := checkRole(r); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	role := make(orm.Params)

	if len(r.Name) > 0 {
		role["Name"] = r.Name
	}
	if len(r.Key) > 0 {
		role["Key"] = r.Key
	}
	if len(r.Description) > 0 {
		role["Description"] = r.Description
	}
	if r.Status != 0 {
		role["Status"] = r.Status
	}
	if len(role) == 0 {
		return 0, errors.New("Update field is empty")
	}
	var table Role
	num, err := o.QueryTable(table).Filter("Id", r.Id).Update(role)
	if err != nil {
		return 0, err
	}
	return num, err
}
func UpdateUserRole(roleId int64, userId int64) (int64, error) {
	o := orm.NewOrm()
	role := Role{Id: roleId}
	user := User{Id: userId}
	m2m := o.QueryM2M(&user, "Role")
	num, err := m2m.Add(&role)
	// m2m.
	return num, err
}
func DelRoleById(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&Role{Id: Id})
	return status, err
}
func GetResourcesByRoleId(Id int64) (resources []orm.Params, count int64) {
	o := orm.NewOrm()
	resource := new(Resource)
	count, err := o.QueryTable(resource).Filter("Role__Role__Id", Id).Values(&resources)
	if err != nil {
		fmt.Println()
		fmt.Println(Id)
		fmt.Println(err)
	}
	return resources, count
}

func GetAllRole() (roles []orm.Params, count int64) {
	o := orm.NewOrm()
	role := new(Role)
	qs := o.QueryTable(role)
	count, err := qs.Values(&roles)
	if err != nil {
		fmt.Println(err)
	}
	return roles, count
}

func DelGroupResource(roleId int64, groupId int64) error {
	var resources []*Resource
	var resource Resource
	role := Role{Id: roleId}
	o := orm.NewOrm()
	num, err := o.QueryTable(resource).Filter("Group", groupId).RelatedSel().All(&resources)
	if err != nil {
		return nil
	}
	if num < 1 {
		return nil
	}
	for _, n := range resources {
		m2m := o.QueryM2M(n, "Role")
		_, err1 := m2m.Remove(&role)
		if err1 != nil {
			return err1
		}
	}
	return nil
}

func AddRoleResource(roleId int64, resourceId int64) (int64, error) {
	o := orm.NewOrm()
	role := Role{Id: roleId}
	resource := Resource{Id: resourceId}
	m2m := o.QueryM2M(&resource, "Role")
	num, err := m2m.Add(&role)
	return num, err
}

func GetRoleByName(Name string) (err error, role Role) {
	o := orm.NewOrm()
	role = Role{Name: Name}
	err = o.Read(&role, "Name")
	return err, role
}
func DelUserRole(roleId int64) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("user_roles").Filter("role_id", roleId).Delete()
	return err
}

func DelUserRoleByUserId(userId int64) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("user_roles").Filter("user_id", userId).Delete()
	return err
}
func AddRoleUser(roleId int64, userId int64) (int64, error) {
	o := orm.NewOrm()
	role := Role{Id: roleId}
	user := User{Id: userId}
	m2m := o.QueryM2M(&user, "Role")
	num, err := m2m.Add(&role)
	return num, err
}

func AccessList(uid int64) (list []orm.Params, err error) {
	var roles []orm.Params
	o := orm.NewOrm()
	role := new(Role)
	_, err = o.QueryTable(role).Filter("User__User__Id", uid).Values(&roles)
	if err != nil {
		return nil, err
	}
	var resources []orm.Params
	resource := new(Resource)
	for _, r := range roles {
		_, err := o.QueryTable(resource).Filter("Role__Role__Id", r["Id"]).Values(&resources)
		if err != nil {
			return nil, err
		}
		for _, n := range resources {
			list = append(list, n)
		}
	}
	return list, nil
}
