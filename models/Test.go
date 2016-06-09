package models

// import (
// 	// "errors"
// 	// "fmt"
// 	"github.com/astaxie/beego"
// 	"github.com/astaxie/beego/orm"
// 	// "github.com/astaxie/beego/validation"

// 	"time"
// )

// // type Test struct {
// // 	// Mesid     int64
// // 	Name      string
// // 	User      string
// // 	Addtime   time.Time `orm:"type(datetime);auto_now_add"`
// // 	Status    int64
// // 	Statuszhi string
// // }

// func (r *Test) TableName() string {
// 	return beego.AppConfig.String("Test_table")
// }

// /**注册model*/
// func init() {
// 	orm.RegisterModel(new(Test))
// }

// func Leftjoin(page int64, pageSize int64) (list []orm.Params) {
// 	o := orm.NewOrm()
// 	sql := "select respond.mesid,respond.user,messsage.statuszhi,messsage.name,messsage.status,messsage.addtime as addtime  from respond left join messsage on respond.mesid=messsage.id order by respond.mesid  asc limit ?,?"
// 	o.Raw(sql, (page-1)*pageSize, pageSize).Values(&list)
// 	fmt.Println("----------------------")
// 	fmt.Println(list)
// 	return list
// }

// func DelMessageByMesid() {
// 	o := orm.NewOrm()
// 	status, err := o.Delete(&list{Mesid: mesid})
// 	return status, err
// }

// func DelMessageByMesid(Mesid int64) (int64, error) {
// 	o := orm.NewOrm()
// 	status, err := o.Delete(&Test{Mesid: Mesid})
// 	return status, err
// }

// func AddTest(r *Test) (int64, error) {
// 	o := orm.NewOrm()
// 	id, err := o.Insert(r)
// 	return id, err
// }
//数据插入尝试insert into test(mesid,user,name,status,addtime)
