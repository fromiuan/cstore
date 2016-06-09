package models

import (
	"errors"
	"fmt"
	"log"
	"time"

	. "cstore/src/tool"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

/*
*用户表
* beego 中会把名为Id的字段自动设置文自增加的主键
 */
type User struct {
	Id       int64
	Uname    string `orm:"unique;size(64)" form:"Uname" valid:"Required"`
	Pwd      string `orm:"size(64)" form:"Pwd"  valid:"Required"`
	Repwd    string `orm:"-" form:"Repwd" valid:"Required" `
	Nickname string `orm:"size(64)" form:"Nickname" valid:"Required"`
	Email    string `orm:"size(64)" form:"Email" valid:"Required;Email"`
	Phone    string `orm :"size(64)" form:"Phone" valid:"Required;Mobile"`
	// Status  2 启用  1禁止
	Status     int64  `orm:"size(11);default(2)" form:"Status" valid:"Required;Range(1,2)"`
	Statusname string `orm:"-" form:"Statusname" valid:"Required" `
	Remark     string `orm:"null;size(200)" form:"Remark" valid:"MaxSize(200)"`
	//角色名称
	Rolename string `orm:"-" form:"Rolename" valid:"Required" `

	Logintime time.Time `orm:"null;type(datetime)" form:"-"`
	Ctime     time.Time `orm:"type(datetime);auto_now_add"`
	Role      []*Role   `orm:"rel(m2m)"`
}

func (u *User) TableName() string {
	return beego.AppConfig.String("user_table")
}

func (u *User) Valid(v *validation.Validation) {
	if u.Pwd != u.Repwd {
		v.SetError("Repwd", "二次输入的密码不一致")
	}
}

func checkUser(u *User) (err error) {
	valid := validation.Validation{}
	b, _ := valid.Valid(&u)
	if !b {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			return errors.New(err.Message)
		}
	}
	return nil
}

func init() {
	orm.RegisterModel(new(User))
}

/*
*获取用户列表
 */
func GetUserList(page int64, page_size int64, sort string) (users []orm.Params, count int64) {
	omodel := orm.NewOrm()
	user := new(User)
	qs := omodel.QueryTable(user)
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}
	qs.Limit(page_size, offset).OrderBy(sort).Values(&users)
	count, _ = qs.Count()
	return users, count

}

func IsExitUser(Id int64) (error, User) {
	o := orm.NewOrm()
	user := User{Id: Id}
	err := o.Read(&user, "Id")
	return err, user
}

/*
*添加用户
 */
func AddUser(u *User) (int64, error) {
	if err := checkUser(u); err != nil {
		return 0, err
	}
	omodel := orm.NewOrm()
	user := new(User)
	user.Uname = u.Uname
	user.Pwd = EncodeUserPwd(u.Uname, u.Pwd)
	user.Nickname = u.Nickname
	user.Email = u.Email
	user.Phone = u.Phone
	user.Status = u.Status
	user.Remark = u.Remark

	id, err := omodel.Insert(user)
	return id, err
}

/*
*更新用户
**/
func UpdateUser(u *User) (int64, error) {
	if err := checkUser(u); err != nil {
		return 0, err
	}
	omodel := orm.NewOrm()
	user := make(orm.Params)
	fmt.Printf("uname is %s\n", u.Uname)
	if len(u.Uname) > 0 {
		user["Uname"] = u.Uname
	}
	if len(u.Pwd) > 0 {
		user["Pwd"] = EncodeUserPwd(u.Uname, u.Pwd)
	}
	if len(u.Nickname) > 0 {
		user["Nickname"] = u.Nickname
	}
	if len(u.Email) > 0 {
		user["Email"] = u.Email
	}
	if len(u.Phone) > 0 {
		user["Phone"] = u.Phone
	}
	if len(u.Remark) > 0 {
		user["Remark"] = u.Remark
	}
	if u.Status != 0 {
		user["Status"] = u.Status
	}
	if len(user) == 0 {
		return 0, errors.New("update field is empty")
	}

	var table User
	num, err := omodel.QueryTable(table).Filter("Id", u.Id).Update(user)
	return num, err

}

func DelUserById(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&User{Id: Id})
	return status, err
}

func GetRoleByUserId(userId int64) (roles Role) {
	fmt.Printf("user id  is %d", userId)
	o := orm.NewOrm()
	role := new(Role)
	err := o.QueryTable(role).Filter("User__User__Id", userId).One(&roles)
	if err != nil {
		fmt.Println(err)
	}
	return roles
}

func GetUserByUname(uname string) (user User) {
	o := orm.NewOrm()
	user = User{Uname: uname}
	err := o.Read(&user, "Uname")
	if err == orm.ErrNoRows {
		fmt.Println("\n 查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("\n 找不到主键")
	} else {
		fmt.Println(err)
		fmt.Println("\n")
		fmt.Println(user.Id, user.Uname, user.Email)
	}
	return user
}

func Users() (cout int64, users []orm.Params) {
	o := orm.NewOrm()
	user := new(User)
	qs := o.QueryTable(user)
	count, _ := qs.Values(&users, "id", "uname", "pwd", "email")
	return count, users
}
