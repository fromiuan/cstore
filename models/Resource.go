package models

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

type Resource struct {
	Id    int64
	Name  string `orm:"size(32)" form:"Name" valid:"Required" `
	Pid   int64  `orm:size(11) form:"Pid" valid:"Required"`
	Pname string `orm:"-" form:"Pname" valid:"Required" `
	Key   string `orm:"size(64)" form :"Key" valid:"Required"`
	Type  string `orm:"size(10)" form:"Type" `
	// 目录 菜单  按钮
	Typename    string  `orm:"-" form:"Typename" valid:"Required" `
	Url         string  `orm:"size(64)" form:"Url"`
	Level       int64   `orm:"default(1);size(11)" form:"Level"`
	Description string  `orm:"null;size(200)" form:"Description" valid:"MaxSize(200)"`
	Group       *Group  `orm:"rel(fk)"`
	Role        []*Role `orm:"rel(m2m)"`
}

func (r *Resource) TableName() string {
	return beego.AppConfig.String("resource_table")
}

func checkResource(r *Resource) (err error) {
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
func init() {
	orm.RegisterModel(new(Resource))
}
func GetResources(page int64, page_size int64, sort string) (resources []orm.Params, count int64) {
	fmt.Printf("page是 %d and  page_size是 %d", page, page_size)
	o := orm.NewOrm()
	resource := new(Resource)
	qs := o.QueryTable(resource)
	var offset int64
	if page < 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}
	fmt.Printf("\npage_size is %d page is %d\n", page_size, page)
	qs.Limit(page_size, offset).OrderBy(sort).Values(&resources)
	count, err := qs.Count()
	if err != nil {
		fmt.Print("err is ")
		fmt.Println(err)
	}
	for _, resource := range resources {
		res, _ := GetResource(resource["Pid"].(int64))
		resource["Pname"] = res.Name
	}
	return resources, count
}

func GetAllResource() (count int64, resources []orm.Params) {
	o := orm.NewOrm()
	resource := Resource{}
	qs := o.QueryTable(resource)
	count, err := qs.Values(&resources)
	if err != nil {
		fmt.Println(err)
	}
	return count, resources
}

func GetParentResource(pid int64) (count int64, resources []orm.Params) {
	o := orm.NewOrm()
	resource := new(Resource)
	qs := o.QueryTable(resource).Filter("Pid", pid)
	count, err := qs.Values(&resources)
	if err != nil {
		fmt.Println(err)
	}
	return count, resources
}

func GetResourceByName(name string) (err error, resource Resource) {
	o := orm.NewOrm()
	resource = Resource{Name: name}
	err = o.Read(&resource, "Name")
	return err, resource
}

func GetResource(id int64) (Resource, error) {
	o := orm.NewOrm()
	resource := Resource{Id: id}
	err := o.Read(&resource)
	if err != nil {
		return resource, nil
	}
	return resource, nil
}
func AddResource(r *Resource) (int64, error) {
	if err := checkResource(r); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	resource := new(Resource)
	resource.Name = r.Name
	resource.Pid = r.Pid
	resource.Key = r.Key
	resource.Type = r.Type
	resource.Url = r.Url
	resource.Level = r.Level
	resource.Description = r.Description
	resource.Group = r.Group

	id, err := o.Insert(resource)
	return id, err
}

func UpdateResource(r *Resource) (int64, error) {
	if err := checkResource(r); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	resource := make(orm.Params)
	if len(r.Name) > 0 {
		resource["Name"] = r.Name
	}
	// if r.Pid != 0 {
	// }
	resource["Pid"] = r.Pid

	if len(r.Key) > 0 {
		resource["Key"] = r.Key
	}
	if len(r.Description) > 0 {
		resource["Description"] = r.Description
	}
	if len(r.Type) > 0 {
		resource["Type"] = r.Type
	}
	if len(r.Url) > 0 {
		resource["Url"] = r.Url
	}

	if r.Level != 0 {
		resource["Level"] = r.Level
	}
	var table Resource
	num, err := o.QueryTable(table).Filter("Id", r.Id).Update(resource)
	if err != nil {
		fmt.Println("err is ", err)
	}
	fmt.Println("num is ", num)
	return num, err
}
func DelResourceById(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&Resource{Id: Id})
	return status, err
}

func DelRoleRescoursByRoleId(roleId int64) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("resource_roles").Filter("role_id", roleId).Delete()
	return err

}
func AddRoleRescours(roleId int64, resList []string) error {
	o := orm.NewOrm()
	role := Role{Id: roleId}
	for _, v := range resList {
		resourceId, _ := strconv.Atoi(v)
		resource := Resource{Id: int64(resourceId)}
		m2m := o.QueryM2M(&resource, "Role")
		m2m.Add(&role)
	}

	return nil
}

func GetResourcesByGroupId(GroupId int64) (resources []orm.Params, count int64) {
	o := orm.NewOrm()
	resource := new(Resource)
	count, _ = o.QueryTable(resource).Filter("Group", GroupId).Values(&resources)
	return resources, count
}

func GetResourceByRoleId(id int64) (resources []orm.Params, count int64) {
	o := orm.NewOrm()
	resource := new(Resource)
	count, _ = o.QueryTable(resource).Filter("Role__Role__Id", id).Values(&resources)
	return resources, count
}

func GetResourceTree(pid int64, resType int64) ([]orm.Params, error) {
	o := orm.NewOrm()
	resource := new(Resource)
	var resources []orm.Params
	_, err := o.QueryTable(resource).Filter("Pid", pid).Filter("Type", resType).OrderBy("Level").Values(&resources)
	if err != nil {
		return resources, err
	}
	return resources, nil
}
func GetResSubTree() {

}
