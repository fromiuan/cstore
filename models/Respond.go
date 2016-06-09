package models

import (
	// "errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	// "github.com/astaxie/beego/validation"

	//"time"
)

type Respond struct {
	Mesid int64 `orm:"pk"`
	Id    int64
	User  string
}

func (r *Respond) TableName() string {
	return beego.AppConfig.String("Respond_table")
}

/**注册model*/
func init() {
	orm.RegisterModel(new(Respond))
}

func AddRespond(r *Respond) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(r)
	return id, err
}

func GetResponById(Id int64) (err error, respond Respond) {
	o := orm.NewOrm()
	respond = Respond{Id: Id}
	err = o.Read(&respond, "Id")
	return err, respond
}

func Leftjoin(page int64, pageSize int64) (list []orm.Params) {
	o := orm.NewOrm()
	sql := "select respond.mesid,respond.user,messsage.statuszhi,messsage.name,messsage.status,messsage.addtime as addtime  from respond left join messsage on respond.mesid=messsage.id where messsage.status=2 order by respond.mesid  asc limit ?,?"
	o.Raw(sql, (page-1)*pageSize, pageSize).Values(&list)
	fmt.Println("----------------------")
	fmt.Println(list)
	return list
}

func Privatemes(page int64, pageSize int64, user string) (list []orm.Params) {
	o := orm.NewOrm()
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println(user)
	sql := "select respond.mesid,respond.user,messsage.statuszhi,messsage.name,messsage.status,messsage.addtime as addtime  from respond left join messsage on respond.mesid=messsage.id where messsage.status='1' and messsage.user=? order by respond.mesid  asc limit ?,?"
	o.Raw(sql, user, (page-1)*pageSize, pageSize).Values(&list)
	fmt.Println("----------------------")
	fmt.Println(list)
	return list
}

// func GetAllPrivate() (list []orm.Params, count int64) {
// 	o := orm.NewOrm()
// 	sql := "select respond.mesid,respond.user,messsage.statuszhi,messsage.name,messsage.status,messsage.addtime as addtime  from respond left join messsage on respond.mesid=messsage.id where messsage.status=2 order by respond.mesid"
// 	o.Raw(sql).Values(&list)
// 	qs := o.QueryTable(list)
// 	count, err := qs.Values(&list)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(count)
// 	return list, count
// }

// func GetAllMassage() (messages []orm.Params, count int64) {
// 	o := orm.NewOrm()
// 	message := new(Message)
// 	qs := o.QueryTable(message)
// 	count, err := qs.Values(&messages)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return messages, count
// }

// func GetUserList(page int64, page_size int64, sort string) (users []orm.Params, count int64) {
// 	omodel := orm.NewOrm()
// 	user := new(User)
// 	qs := omodel.QueryTable(user)
// 	var offset int64
// 	if page <= 1 {
// 		offset = 0
// 	} else {
// 		offset = (page - 1) * page_size
// 	}
// 	qs.Limit(page_size, offset).OrderBy(sort).Values(&users)
// 	count, _ = qs.Count()
// 	return users, count

// }

func DelMessageByMesid(mesid int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&Respond{Mesid: mesid})
	return status, err
}

// func DelMessageById(Id int64) (int64, error) {
// 	o := orm.NewOrm()
// 	status, err := o.Delete(&Message{Id: Id})
// 	return status, err
// }

// func DelMessageById(Id int64) (int64, error) {
// 	o := orm.NewOrm()
// 	status, err := o.Delete(&Message{Id: Id})
// 	return status, err
// }
